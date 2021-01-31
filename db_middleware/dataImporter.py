from models import *
import csv

areas = ["A","B","C"]
for a in areas:
    area = Area()
    area.name = a
    db.session.add(area)
db.session.commit()

with open('seatShow.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    line_count = 0
    for row in csv_reader:
        if line_count == 0:
            print(f'Column names are {", ".join(row)}')
        else:
            seat = Seat()
            seat.seatCode = row[0]
            seat.aid = int(row[1])
            db.session.add(seat)
            print(f"seatCode:{row[0]}, aid: {row[1]}")
        line_count += 1
    db.session.commit()
    print(f'Processed {line_count} lines.')

with open('ntList.csv') as csv_file2:
    csv_reader2 = csv.reader(csv_file2, delimiter=',')
    for row in csv_reader2:
        user = User()
        user.username = row[1]
        user.name = row[0]
        user.password = "Passw0rd!"
        print(f"name:{row[0]}, username: {row[1]}")
        db.session.add(user)
    db.session.commit()