#encoding:utf8

import psycopg2 as pg
import csv
import os, re


SAVE_DIR = "/bootstrap/save"

def get_files(path, hot="hot", y="y"):
    result = dict(hot=dict())
    for path, dire, filename in os.walk(SAVE_DIR):
        fullpath = os.path.join(path, filename)
        if dire == hot:
            r = result[hot]
            if re.match("^MEDIS\d{8}.TXT$", filename):
                r["main"] = fullpath
            elif re.match("^\d{8}.txt", filename):
                r["extra"] = r.get("extra", []).append(fullpath)
            elif re.match("^h\d{8}del.txt", filename):
                r["delete"] = fullpath
        if dire == y:
            if filename == "y.csv":
                result[y] = fullpath


def main():
    infiles = get_files(SAVE_DIR)
    print(infiles)



if __name__ == '__main__':
    main()
