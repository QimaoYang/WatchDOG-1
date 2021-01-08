'''
This file contains code related to user information, including the following functions
'''
from flask_restplus import Resource, Namespace
from models import *
from flask_jwt_extended import get_jwt_identity, jwt_required
import json
from flask import request
from forms import *
from sqlalchemy import or_


def jsontifySeat(seat):
    return {
        "seatCode": seat.seatCode,
        "aid": seat.aid,
        "type": seat.type
    }

def jsontifySeatWithList(seat,l):
    if seat.id in l:
        return {
            f"{seat.seatCode}": "not availble"
        }
    else:
        return {
            f"{seat.seatCode}": "availble",
        }


api = Namespace('', description='Reservation')

@api.route('/v1/db/seat')
class getAvailable(Resource):
    @api.doc(description="Get all available seats")
    @api.param(name='region', _in='url')
    def get(self):
        region = request.args.get('region')
        if region:
            seatcode = db.session.query(Area).filter(Area.name == region).first()
            if seatcode:
                aid = seatcode.id
                resv = db.session.query(Reservation.seat_id).filter(
                    Reservation.release_time >= datetime.now()).all()
                query = db.session.query(Seat).filter(Seat.aid == aid).all()
                return {
                    "data":{
                        "seat": [jsontifySeatWithList(s, [x[0] for x in resv]) for s in query]
                    }
                }, 200
            else:
                return 400
        else:
            seatcode = db.session.query(Area).all()
            print()
            result = []
            for r in seatcode:
                tmp = {}
                aid = r.id
                resv = db.session.query(Reservation.seat_id).filter(
                    Reservation.release_time >= datetime.now()).all()
                query = db.session.query(Seat).filter(Seat.aid == aid).filter(
                    Seat.id.notin_([x[0] for x in resv])).count()
                tmp[r.name] = query
                result.append(tmp)
            return {
                    "data":{
                        "seat": result
                    }
                }, 200


# @api.route('/available')
# class getAvailable(Resource):
#     @api.doc(description="Get all available seats")
#     def get(self):
#         resv = db.session.query(Reservation.seat_id).filter(Reservation.date == date.today()).filter(Reservation.release_time >= datetime.now()).all()
#         resv = [x[0] for x in resv]
#         print(resv)
#         query = db.session.query(Seat).filter(Seat.id.notin_(resv)).filter(Seat.team).all()
#         return [jsontifySeat(s) for s in query], 200
#
# @api.route('/available/<int:aid>')
# class getAvailable(Resource):
#     @api.doc(description="Get all available seats filter by area")
#     def get(self, aid):
#         resv = db.session.query(Reservation.seat_id).filter(Reservation.date == date.today()).filter(Reservation.release_time >= datetime.now()).all()
#         query = db.session.query(Seat).filter(Seat.aid == aid).filter(Seat.id.notin_([x[0] for x in resv])).all()
#         return [jsontifySeat(s) for s in query], 200
#
#
# @api.route('/available/<int:aid>/count')
# class getAvailable(Resource):
#     @api.doc(description="Get all available seat count filter by area")
#     def get(self, aid):
#         resv = db.session.query(Reservation.seat_id).filter(Reservation.date == date.today()).filter(Reservation.release_time >= datetime.now()).all()
#         query = db.session.query(Seat).filter(Seat.aid == aid).filter(Seat.id.notin_([x[0] for x in resv])).count()
#         available = db.session.query(Seat).filter(Seat.aid == aid).count()
#         return {
#                    "max": query,
#                    "available": available
#                }, 200


@api.route('/v1/seat/release')
class getAvailable(Resource):
    @api.doc(description="Release a reservation")
    @jwt_required
    @api.param("Authorization", _in='header')
    def post(self):
        try:
            current_user = get_jwt_identity()
            user = User.query.filter_by(username=current_user).first()

            exists_result = db.session.query(Reservation).with_lockmode("update").filter(Reservation.user_id == user.id).filter(Reservation.release_time >= datetime.now()).first()
            if exists_result:
                exists_result.release_time = datetime.now()
                #exists_result.end_time = datetime.now()
                db.session.commit()
                db.session.refresh(exists_result)
                return exists_result.id, 200
            else:
                return {"message": "The didn't been reserved"}, 400
        except:
            return {"message": "bad payload"}, 400


resource_fields = api.model('reservations_form', reservations_form)
@api.route('/v1/db/seat/register')
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

            s = db.session.query(Seat).filter(Seat.seatCode == r["seat_code"]).filter(or_(Seat.team == "public", Seat.team == user.team)).first()
            #s = db.session.query(Seat).filter(Seat.seatCode == r["seat_code"]).filter(Seat.team == user.team).first()
            if not s:
                return {"message": "No such seat"}, 400
            reservations.seat_id = s.id

            exists_result = db.session.query(Reservation).with_lockmode("update").filter(Reservation.seat_id == s.id).filter(Reservation.release_time >= datetime.now()).order_by(Reservation.created).first()
            if not exists_result:
                db.session.add(reservations)
                db.session.commit()
                db.session.refresh(reservations)
                return reservations.id, 200
            else:
                return {"message": "The seat has been reserved"}, 400
        except:
           return {"message": "bad payload"}, 400

@api.route('/v1/db/user/seat')
class getAvailable(Resource):
    @api.doc(description="get seat code")
    @jwt_required
    @api.param("Authorization", _in='header')
    def get(self):
        try:
            current_user = get_jwt_identity()
            user = User.query.filter_by(username=current_user).first()

            exists_result = db.session.query(Reservation).with_lockmode("update").filter(
                Reservation.user_id == user.id).filter(Reservation.release_time >= datetime.now()).first()
            if exists_result:
                s = db.session.query(Seat).filter(Seat.id == exists_result.seat_id).first()
                return {
                           "seat": s.seatCode
                       }, 200
            else:
                return {
                           "seat": None
                       }, 200
        except:
            return {"message": "bad payload"}, 400

