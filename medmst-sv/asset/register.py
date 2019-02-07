#encoding:utf8

import psycopg2 as pg
import csv
import os, re
import sys
import os.path 
import codecs

ASSET_DIR = "/asset"

if "MEDMST_SAVE" in os.environ:
    SAVE_DIR = os.environ["MEDMST_SAVE"]
else:
    SAVE_DIR = "/bootstrap/save"

PARAM = dict(
            host = os.environ["PG_HOST"],
            port = os.environ["PG_PORT"],
            user = os.environ["PG_USER"],
            password = os.environ["PG_PASSWORD"],
            database = os.environ["PG_DATABASE"],
)

def connection():
    return _connection(PARAM)

def _connection(param):
    return pg.connect(**param)

def get_files(save_dir=SAVE_DIR):
    hot = "hot"
    y = "y"
    medis = "medis"
    result = dict(medis=[], y=[])
    for path, dirs, filenames in os.walk(save_dir):
        for filename in filenames:
            fullpath = os.path.join(path, filename)
            if os.path.basename(path) == hot:
                r = result[medis]
                if re.match("^MEDIS\d{8}.TXT$", filename):
                    r.append(fullpath)
                elif re.match("^\d{8}.txt", filename):
                    r.append(fullpath)
                elif re.match("^h\d{8}del.txt", filename):
                    r.append(fullpath)
            if os.path.basename(path) == y:
                if filename == "y.csv":
                    result[y] = [fullpath]

    print("reuslt", result)
    return result

def _sql_from_file(filepath):
    with codecs.open(filepath, "r", "utf8") as f:
        lines = [line for line in f]
        sql = "\n".join(lines)
    return sql

def create(con):
    filepath = os.path.join(ASSET_DIR, "medis_def.txt")
    sql = _sql_from_file(filepath)
    cur = con.cursor()
    try:
        cur.execute(sql)
    except Exception, e:
        print e

    filepath = os.path.join(ASSET_DIR, "y_def.txt")
    sql = _sql_from_file(filepath)
    cur = con.cursor()
    try:
        cur.execute(sql)
    except Exception, e:
        print e

def insert(con, infiles):
    infiles = get_files()
    
    insert_list = []
    for (table, skip) in [("medis", True), ("y", False)]:
        sql_template = os.path.join(SAVE_DIR, "{0}_insert.txt").format(table)
        insert_data =  infiles[table]
        insert_list.append([sql_template, insert_data, skip])
    for (sql_template, insert_data, skip) in insert_list:
        _insert(con, sql_template, insert_data, skip)

def _insert(con, sql_file, insert_files, line1skip):
    try:
        sql = _sql_from_file(sql_file)
    except IOError, e:
        print(e)
        return
    for insert_file in insert_files:
        with codecs.open(insert_file, "r", "utf-8") as f:
            r = csv.reader(f)
            if line1skip:
                r.next()
            cur = con.cursor()
            cur.executemany(sql, r)

def delete(con):
    sqls = [
        """DELETE FROM "medis";""",
        """DELETE FROM "y";""",
    ]
    cur = con.cursor()
    for sql in sqls:
        cur.execute(sql)

def C():
    con = connection()
    create(con)
    con.commit()

def I():
    con = connection()
    infiles = get_files()
    insert(con, infiles)
    con.commit()

def D():
    con = connection()
    delete(con)
    con.commit()


def main():
    infiles = get_files(SAVE_DIR)

    options = sys.argv[1].lstrip("-")

    
    if len(options) == 0:
        print("OPTION must be -[C][D][I]")
        print("C: create table")
        print("D: delete table data")
        print("I: insert data to table")
        return
    
    #OPTIONの分だけcreate delete insert関数登録
    exec_list = []
    if "C" in options:
        exec_list.append(C)
    if "I" in options:
        exec_list.append(I)
    if "D" in options:
        exec_list.append(D)

    for func in exec_list:
        func()

if __name__ == '__main__':
    main()