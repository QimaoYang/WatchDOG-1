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
from flask_jwt_extended import  create_access_token, create_refresh_token, get_jti, get_jwt_identity, jwt_required, get_raw_jwt

# set ACCESS TOKEN EXPIRES TIME
ACCESS_EXPIRES = timedelta(days=1)
REFRESH_EXPIRES = timedelta(days=30)


# init api namespace
api = Namespace('Auth', description='Auth operations')
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
        refresh_token = create_refresh_token(identity=json.loads(request_body)["username"])

        # check token
        access_jti = get_jti(encoded_token=access_token)
        refresh_jti = get_jti(encoded_token=refresh_token)

        # return tokens
        return {
                   'access_token': access_token,
                   'refresh_token': refresh_token,
                   'uid': user.id
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
            password = json.loads(request_body)["password"]
            name = json.loads(request_body)["name"]
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
        new_user.name = name

        # gen token
        access_token = create_access_token(identity=json.loads(request_body)["username"])
        refresh_token = create_refresh_token(identity=json.loads(request_body)["username"])
        #access_jti = get_jti(encoded_token=access_token)
        #refresh_jti = get_jti(encoded_token=refresh_token)
        try:
            # commit new user to database
            db.session.add(new_user)
            db.session.commit()
        except:
            return {"message": "cannot connect to db"}, 400
        return {'access_token': access_token,
                'refresh_token': refresh_token}, 201

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
        current_user = get_jwt_identity()
        # get user info by token
        user = User.query.filter_by(username=current_user).first()
        request_body = request.data.decode()
        request_body = json.loads(request_body)
        #reset password
        user.password = request_body['password']
        #commit to database
        db.session.commit()
        return 200

#bind api to routes
api.add_resource(Reg, '/signup')
api.add_resource(Login, '/login')
api.add_resource(LogOut, '/logout')
api.add_resource(changePassword, '/password')