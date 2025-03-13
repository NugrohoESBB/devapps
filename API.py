import random
import requests
import datetime
import names
import time

url = "http://192.168.100.68:9926/addLogData"

while True:
    dateTime = datetime.datetime.now()
    D = dateTime.strftime("%Y-%m-%d")
    T = dateTime.strftime("%X")

    N = names.get_first_name()
    L = names.get_last_name()
    K = round(random.uniform(0.0, 100.0), 2)
    I = round(random.uniform(0.0, 100.0), 2)
    F = round(random.uniform(0.0, 100.0), 2)
    A = round(random.uniform(0.0, 100.0), 2)

    data = {
        "d": D,
        "t": T,
        "n": N,
        "l": L,
        "k": K,
        "i": I,
        "f": F,
        "a": A,
    }

    print(f"\n[{D} {T}] API: {data}")

    try:
        response = requests.post(url, json=data, timeout=10)
        print("Status Code:", response.status_code)
        print("Response JSON:", response.json() if response.headers.get('Content-Type') == 'application/json' else response.text)
    except requests.exceptions.SSLError:
        print("SSL Error: Pastikan server memiliki sertifikat valid atau gunakan HTTP.")
    except requests.exceptions.ConnectionError:
        print("Connection Error: Tidak bisa terhubung ke server. Pastikan IP benar dan server aktif.")
    except requests.exceptions.Timeout:
        print("Timeout Error: Server tidak merespons dalam waktu yang ditentukan.")
    except requests.exceptions.RequestException as e:
        print("Request Error:", e)

    print("Waiting...\n")
    try:
        time.sleep(360)
    except KeyboardInterrupt:
        print("\nByee...")
        break
