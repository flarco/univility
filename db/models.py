import datetime

from sqlalchemy import (
  create_engine,
  Table,
  Column,
  MetaData,
  String,
  Date,
  DateTime,
  Numeric,
  Integer,
)

from sqlalchemy.sql import text

engine = create_engine('sqlite:///test.db', echo=True)


metadata = MetaData()
jobs = Table('jobs', metadata,
  Column('job_id', Integer, primary_key=True, autoincrement=True),
  Column('name', String),
  Column('command', String),
  Column('arguments', String),
  Column('created_at', DateTime, default=datetime.datetime.now),
  Column('updated_at', DateTime, default=datetime.datetime.now),
)

schedules = Table('schedules', metadata,
  Column('scehdule_id', Integer, primary_key=True, autoincrement=True),
  Column('job_id', Numeric),
  Column('queue', String),
  Column('start_time', DateTime),
  Column('end_time', DateTime),
  Column('status', String),
  Column('pid', Numeric),
  Column('return_code', Numeric),
  Column('output', String),
  Column('created_at', DateTime, default=datetime.datetime.now),
  Column('updated_at', DateTime, default=datetime.datetime.now),
)

metadata.drop_all(engine)
metadata.create_all(engine)
conn = engine.connect()

# get_seq = lambda t: conn.execute(
# sequences = {t:get_seq(t) for t in metadata.tables.keys()}


rec = dict(
  name='Test Job',
  command='python /path/to/script.py',
  arguments='''hello arg1 "select 'hello' from Table"''',
)

sql = jobs.insert().values(**rec)
conn.execute(sql)
res = conn.execute(sql)

# sql = jobs.select().order_by('-job_id')
# res = conn.execute(sql).first()


print(res.inserted_primary_key[0])



