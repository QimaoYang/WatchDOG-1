'''
This file contains the code for the authentication function and contains the following functions
    *log in
    *log out
    *registered
    *change pass word
'''
from flask import request, jsonify
import json
from flask_restplus import Resource, Namespace
from forms import *
from datetime import timedelta
from models import *
from flask_jwt_extended import create_access_token, create_refresh_token, get_jti, get_jwt_identity, jwt_required, get_raw_jwt,jwt_refresh_token_required
from app import blacklist
import csv

NT_LIST = []
with open('ntList.csv') as csv_file2:
    csv_reader2 = csv.reader(csv_file2, delimiter=',')
    for row in csv_reader2:
        NT_LIST.append(row[1])

# init api namespace
api = Namespace('', description='Auth operations')
resource_fields = api.model('login_form', login_form)


#login api
class Login(Resource):
    @api.expect(resource_fields)
    @api.doc(description="login operation, will return a access token and a refresh token")
    def post(self):
        # decode request body
        try:
            request_body = request.data.decode()
            username = json.loads(request_body)["username"]
            password = json.loads(request_body)["password"]
        except:
            return {"message": "input"}, 400
        # get user entity
        user = User.query.filter_by(username=username).first()
        # check if user in database
        if not user:
            return {"message": "wrong username"}, 400
        elif user and not user.check_password(password):
            return {"message": "wrong password"}, 400
        # gen token
        access_token = create_access_token(identity=json.loads(request_body)["username"])


        # return tokens
        return {
                   'session_key': access_token
               }, 200


# Bind input form
reg_forms = api.model('reg_form', reg_form)

class Reg(Resource):
    @api.expect(reg_forms)
    @api.doc(description="signon operation, will return a access token and a refresh token")
    def post(self):
        # decode request body
        try:
            request_body = request.data.decode()
            print(json.loads(request_body))
            username = json.loads(request_body)["username"]

            if username not in NT_LIST:
                return {"message": "username not in nt_list"}, 400
            password = json.loads(request_body)["password"]
            if len(password) < 6:
                return {"message": "password too short"}, 400
        except:
            return {"message": "wrong payload"}, 400
        # check if username already in database(it should be unique)
        user = User.query.filter_by(username=username).first()
        if user:
            return {"message": "Username has been used"}, 400

        # gen new user
        new_user = User()
        new_user.username = username
        new_user.password = password

        # gen token
        access_token = create_access_token(identity=json.loads(request_body)["username"])

        try:
            # commit new user to database
            db.session.add(new_user)
            db.session.commit()
        except:
            return {"message": "cannot connect to db"}, 400
        return {'session_key': access_token}, 201

#log out function
class LogOut(Resource):
    @jwt_required
    @api.doc(description="m")
    @api.param("Authorization", _in='header')
    @api.doc(description="logout operation，revoke the token")
    def delete(self):
        jti = get_raw_jwt()['Authorization']
        #revoke the token
        return jsonify({"msg": "Access token revoked"}), 200


#change password
password_form = api.model('password_form', password_form)
class changePassword(Resource):
    @api.expect(password_form)
    @jwt_required
    @api.param("Authorization", _in='header')
    @api.doc(description="change password")
    def post(self):
        try:
            current_user = get_jwt_identity()
            request_body = request.data.decode()
            request_body = json.loads(request_body)
            password = request_body['password']
            if len(password) < 6:
                return {"message": "password too short"}, 400
            # get user info by token
            user = User.query.filter_by(username=current_user).first()
            #reset password
            user.password = password
            jti = get_raw_jwt()['jti']
            blacklist.add(jti)
            #commit to database
            db.session.commit()
        except:
            return {"message": "bad payload"}, 400
        return 200

reset_form = api.model('password_form', reset_form)

class resetPassword(Resource):
    @api.expect(reset_form)
    @jwt_required
    @api.param("Authorization", _in='header')
    @api.doc(description="reset password")
    def delete(self):
        try:
            current_user = User.query.filter_by(username=get_jwt_identity()).first()
            request_body = request.data.decode()
            request_body = json.loads(request_body)
            username = request_body['username']
            if current_user.username not in ["1259540", "1251337"]:
                return 400
            # get user info by token
            user = User.query.filter_by(username=username).first()
            #delete user
            db.session.delete(user)
            db.session.commit()
        except:
            return {"message": "bad payload"}, 400
        return 200

    @api.expect(reset_form)
    @jwt_required
    @api.param("Authorization", _in='header')
    @api.doc(description="reset password")
    def post(self):
        try:
            current_user = User.query.filter_by(username=get_jwt_identity()).first()
            request_body = request.data.decode()
            request_body = json.loads(request_body)
            username = request_body['username']
            if current_user.username not in ["1259540", "1251337"]:
                return 400
            # get user info by token
            user = User.query.filter_by(username=username).first()
            # reset password
            user.password = "passw0rd!"
            db.session.commit()
        except:
            return {"message": "bad payload"}, 400
        return 200

class refresh(Resource):
    @api.param("Authorization", _in='header')
    @jwt_refresh_token_required
    def post(self):
        current_user = get_jwt_identity()
        ret = {
            'access_token': create_access_token(identity=current_user)
        }
        return ret, 200


#bind api to routes
api.add_resource(Reg, '/v1/db/user/register')
#api.add_resource(Reg, '/signup')
#api.add_resource(Login, '/login')
api.add_resource(Login, '/v1/db/user/login')
#api.add_resource(LogOut, '/logout')
api.add_resource(changePassword, '/v1/db/user/password')
api.add_resource(resetPassword, '/v1/db/user/reset')

#api.add_resource(refresh, '/refresh')