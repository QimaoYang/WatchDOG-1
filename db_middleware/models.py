'''
This file contains the code related to the database structure
'''

from sqlalchemy import Column, Integer, String, Text, DateTime ,ForeignKey, Date
from werkzeug.security import generate_password_hash, check_password_hash
from datetime import datetime
from datetime import date
from app import db
from datetime import timedelta

#db table structure
class User(db.Model):
    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True,autoincrement=True)
    name = db.Column(Text)
    _password = db.Column('password', db.String(128), nullable=False)
    username = db.Column(db.String(512),unique=True)
    team = db.Column(db.String)

    @property
    def password(self):
        return self._password

    @password.setter
    def password(self, value):
        self._password = generate_password_hash(value)

    def check_password(self, user_pwd):
        return check_password_hash(self._password, user_pwd)


class Area(db.Model):
    __tablename__ = 'areas'
    id = db.Column(db.Integer, primary_key=True,autoincrement=True)
    name = db.Column(Text)


class Seat(db.Model):
    __tablename__ = 'seats'
    id = db.Column(db.Integer, primary_key=True,autoincrement=True)
    seatCode = db.Column(Text)
    aid = db.Column(db.Integer, ForeignKey('areas.id'))
    type = db.Column(String(20))
    team = db.Column(db.String, default="public")




class Reservation(db.Model):
    __tablename__ = 'seat_reservations'
    id = db.Column(db.Integer, primary_key=True,autoincrement=True)
    seat_id = db.Column(db.Integer, ForeignKey('seats.id'))
    user_id = db.Column(db.Integer)
    date = db.Column(db.Date, default=date.today())
    start_time = db.Column(db.DateTime, default=datetime.now())
    end_time = db.Column(db.DateTime, default=datetime.now()+timedelta(hours=9))
    created = db.Column(db.DateTime, default=datetime.now())
    release_time = db.Column(db.DateTime, default=datetime.now()+timedelta(hours=9))
    ip = db.Column(Text)
