'''
This file contains the code to generate the database
'''
from sqlalchemy.orm import sessionmaker
from models import *
db.metadata.create_all(db.engine)
DBSession = sessionmaker(bind=db.engine)
session = DBSession()