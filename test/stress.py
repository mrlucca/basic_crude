import requests
from multiprocessing import Pool


def acess_page(n):
    print(requests.get("http://localhost:7000/login"))
    print(requests.get("http://localhost:7000/login/register"))

with Pool(100) as p:
    p.map(acess_page, [x for x in range(30*50000)])
