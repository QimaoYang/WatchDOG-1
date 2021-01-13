'''
This file contains the basic settings of the project, include:
    * database path
'''

import os
basedir = os.path.abspath(os.path.dirname(__file__))
from datetime import timedelta

# settings of backend server
class Config():

    os.environ.get("SECRET_KEY") or "765 PRODUCTION"
    #set database path
    SQLALCHEMY_DATABASE_URI = 'sqlite:///test.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False

    # set token secret key
    JWT_SECRET_KEY = "minatoaqua"
    JWT_REFRESH_TOKEN_EXPIRES = False
    #JWT_EXPIRATION_DELTA = timedelta(hours=24)
    JWT_BLACKLIST_ENABLED = True
    JWT_BLACKLIST_TOKEN_CHECKS = ['access', 'refresh']

