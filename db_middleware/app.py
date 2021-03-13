'''
Main program
This file contains the code for the following functions
    *setting base url
    *set secret key
    *link data base
    *bind sub APIs to swagger
'''
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from config import Config
from flask_restplus import Api
from flask_cors import CORS
from flask_jwt_extended import JWTManager


#set swagger info
api: Api = Api(
    title='kic2',
    version='1.0',
    description='v1.0',
    prefix='/powercubicle'
)
app = Flask(__name__)
#set key
app.config.from_object(Config)
api.init_app(app)

#link database
db = SQLAlchemy(app)
migrate = Migrate(app, db)
jwt = JWTManager(app)

#JWT Black list
blacklist = set()
@jwt.token_in_blacklist_loader
def check_if_token_in_blacklist(decrypted_token):
    jti = decrypted_token['jti']
    return jti in blacklist

#CORS added
CORS(app)


#blind namespace to swagger api page
if __name__ == '__main__':
    app.debug = False

    from authRoutes import api as ns1
    from reservationRoutes import api as ns2
    from reportGenerator import api as ns3

    #bind apis
    api.add_namespace(ns1)
    api.add_namespace(ns2)
    api.add_namespace(ns3)

    #run backend server at port 5001
    app.run(host="0.0.0.0",port=5001)
    from waitress import serve
    #serve(app, host="0.0.0.0", port=5001)