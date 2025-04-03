import requests

def test_proxy():
    proxy = "http://localhost:8080"
    target_url = "https://www.examasdasads.com/"
    
    proxies = {
        "http": proxy,
        "https": proxy,
    }
    
    try:
        response = requests.get(target_url, proxies=proxies)
        print("Response status:", response.status_code)
        print("Response body:", response.text[:500])  # Print first 500 characters
    except Exception as e:
        print("Error:", e)

if __name__ == "__main__":
    test_proxy()