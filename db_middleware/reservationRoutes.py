'''
This file contains code related to user information, including the following functions
'''
from flask_restplus import Resource, Namespace
from models import *
from flask_jwt_extended import get_jwt_identity, jwt_required
import json
from flask import request
from forms import *


def jsontifySeat(seat):
    return {
        "seatCode": seat.seatCode,
        "aid": seat.aid,
        "type": seat.type
    }


api = Namespace('Reservation', description='Reservation')


@api.route('/available')
class getAvailable(Resource):
    @api.doc(description="Get all available seats")
    def get(self):
        resv = db.session.query(Reservation.seat_id).filter(Reservation.date == date.today()).all()
        query = db.session.query(Seat).filter(Seat.id.notin_(resv)).all()
        return [jsontifySeat(s) for s in query], 200


@api.route('/available/<int:aid>')
class getAvailable(Resource):
    @api.doc(description="Get all available seats filter by area")
    def get(self, aid):
        resv = db.session.query(Reservation.seat_id).filter(Reservation.date == date.today()).all()
        query = db.session.query(Seat).filter(Seat.aid == aid).filter(Seat.id.notin_(resv)).all()
        return [jsontifySeat(s) for s in query], 200


@api.route('/available/<int:aid>/count')
class getAvailable(Resource):
    @api.doc(description="Get all available seat count filter by area")
    def get(self, aid):
        resv = db.session.query(Reservation.seat_id).filter(Reservation.date == date.today()).all()
        query = db.session.query(Seat).filter(Seat.aid == aid).filter(Seat.id.notin_(resv)).count()
        available = db.session.query(Seat).filter(Seat.aid == aid).count()
        return {
                   "max": query,
                   "available": available
               }, 200


resource_fields = api.model('reservations_form', reservations_form)


@api.route('/reservations')
class getAvailable(Resource):
    @api.doc(description="Create a reservation")
    @jwt_required
    @api.expect(resource_fields)
    @api.param("Authorization", _in='header')
    def post(self):
        try:
            r = request.data.decode()
            r = json.loads(r)
            current_user = get_jwt_identity()
            user = User.query.filter_by(username=current_user).first()

            # create a reservations
            reservations = Reservation()
            # checklist = ["start_time","end_time","ip","release_time"]

            reservations.user_id = user.id
            reservations.seat_id = r["seat_id"]
            exists_result = db.session.query(Reservation).with_lockmode("update").filter(
                Reservation.seat_id == r["seat_id"]).filter(Reservation.date == date.today()).first()
            if not exists_result:
                db.session.add(reservations)
                db.session.commit()
                db.session.refresh(reservations)
                return reservations.id, 200
            else:
                return {"message": "The seat has been reserved"}, 400
        except:
            return {"message": "bad payload"}, 400