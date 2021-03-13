from models import *
import datetime
from datetime import timedelta
import matplotlib.pyplot as plt
import numpy as np
from flask import Flask, render_template, make_response
from flask_restplus import Resource, Namespace
from io import BytesIO
import base64


def increase_by_key(r, key, t):
    if key in r.keys():
        if t in r[key].keys():
            r[key][t] += 1
        else:
            r[key][t] = 1
        r[key]["total"] += 1
    else:
        r[key] = {t: 1, "total": 1}




api = Namespace('', description='report')

@api.route('/v1/db/report')
class getReport(Resource):
    @api.doc(description="Get booking report")
    def get(self):
        result = {}
        img1 = BytesIO()
        img2 = BytesIO()
        for i in reversed(range(1, 8)):
            res = db.session.query(Reservation).filter(
                Reservation.created > datetime.datetime.now().date() - timedelta(days=i)).filter(
                Reservation.created < datetime.datetime.now().date() - timedelta(days=i - 1)).order_by(
                Reservation.date).all()
            db.session.close()
            total = 0
            for r in res:
                r.created += timedelta(hours=8)
                key = str(r.created.date())[5:]
                if r.created.time() < datetime.time(9, 0, 0):
                    increase_by_key(result, key, "before 9:00")
                elif datetime.time(9, 0, 0) <= r.created.time() < datetime.time(10, 0, 0):
                    increase_by_key(result, key, "9:00~10:00")
                elif datetime.time(10, 0, 0) <= r.created.time() < datetime.time(11, 0, 0):
                    increase_by_key(result, key, "10:00~11:00")
                elif datetime.time(11, 0, 0) <= r.created.time() < datetime.time(12, 0, 0):
                    increase_by_key(result, key, "11:00~12:00")
                elif datetime.time(12, 0, 0) <= r.created.time():
                    increase_by_key(result, key, "after 12:00")

        res = [(k, v["total"]) for k, v in result.items()]
        x = np.array([x[0] for x in res])
        y = np.array([x[1] for x in res])
        for xx, yy in result.items():
            plt.text(xx, yy["total"], str(yy["total"]), ha='center')
        plt.xlabel("Date")
        plt.ylabel("Total")
        plt.title("Booking Count Distribute(Last Week)")
        plt.bar(x, y)
        plt.savefig(img1, format='png')
        plt.close()
        img1.seek(0)
        plot_url1 = base64.b64encode(img1.getvalue()).decode('utf8')

        result2 = {}
        for kk, yy in result.items():
            for k, v in yy.items():
                if k == "total":
                    continue
                if k in result2.keys():
                    result2[k] += v
                else:
                    result2[k] = v
        res = []
        for x in ["before 9:00", "9:00~10:00", "10:00~11:00", "11:00~12:00", "after 12:00"]:
            if x in result2.keys():
                res.append((x, result2[x]))

        x = np.array([x[0] for x in res])
        y = np.array([x[1] for x in res])
        for xx, yy in res:
            plt.text(xx, yy, str(yy), ha='center')
        plt.xlabel("Time")
        plt.ylabel("Total")
        plt.title("Booking Time Distribute(Last Week)")
        plt.bar(x, y)
        plt.savefig(img2, format='png')
        plt.close()
        img2.seek(0)
        plot_url2 = base64.b64encode(img2.getvalue()).decode('utf8')

        resp = make_response(render_template('report.html', plot_url1=plot_url1, plot_url2=plot_url2), 200)
        resp.headers['Content-Type'] = 'text/html; charset=utf-8'
        return resp

