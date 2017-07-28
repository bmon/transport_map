import sqlite3
import plotly.offline as plot
import plotly.graph_objs as go

conn = sqlite3.connect("./cities.db")
cursor = conn.cursor()

points = cursor.execute("select lat, lng, duration, duration_in_traffic from points where status=\"OK\";").fetchall()

rows = [[]]
prev = points[0]
sensitivity = 1000
for cur in points:
    # rebuild lat-lng data into matrix
    if abs(cur[0] - prev[0]) * sensitivity > 1:
        # new row
        rows.append([])
    # add new z value
    rows[-1].append(cur[2])
    prev = cur

trace = go.Heatmap(z = rows[::-1])
print(plot.plot([trace], filename="plotly-heatmap.html", auto_open=False))

