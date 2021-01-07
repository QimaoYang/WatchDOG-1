'''
This file all request form of this project
'''
from flask_restplus import fields
login_form =  {'username': fields.String(default='string: username',example="xx_yy@dell.com", required=True),
        'password': fields.String(default='string: password',required=True)}
reg_form = {#'name': fields.String(default='string: xx_yy', required=True),
        'password': fields.String(default='string: sss',required=True),
        'username': fields.String(default='string: xx_yy@dell.com', example='xx_yy@dell.com',required=True),
        #'team': fields.String(default='string: xx_yy', required=True)
}
password_form = {'password':fields.String(required=True)}

reservations_form = {'seat_code':fields.String(required=True)}
