PROJECT SETUP INSTRUCTIONS:

✅ Prerequisites:

Docker installed 

Docker Compose installed 

Git installed

Internet connection 


📁 Folder Structure:
![WhatsApp Image 2025-06-25 at 9 56 43 AM](https://github.com/user-attachments/assets/a53129a8-5593-45c1-b1ed-d5282ac2e905)


🧪 Step-by-step Setup:

1. Clone the project:

               git clone https://github.com/Ramkumar120/dpdzero-task.git

               cd dpdzero-task

2. Build and Start All Containers:

                sudo docker compose down        # Clean up any previous setup
                sudo docker compose up -d --build 

3. Verify Everything is Running:

                sudo docker ps

Look for 3 containers: nginx, service1, service2
Check for (healthy) in status

🌐 Test the Services:

Test Go app:

                curl http://localhost/service1/hello

Test Python app:

                curl http://localhost/service2/hello

Health Check Endpoints:

                curl http://localhost/service1/health
                curl http://localhost/service2/health

🛑 Stop and Clean:

                sudo docker compose down

📌 Extras:

1. Nginx logs stored in volume: ./nginx-logs-volume

2. Health checks are configured in docker-compose.yml

3. Uses bridge networking only (no host networking)

4. Created health check endpoints /health for both service1 and service2. 

5. Created Jenkins pipeline to automate this entire process . Whenever a code push to the github, jenkins will automatically trigger the pipeline using webhook.
   ![WhatsApp Image 2025-06-25 at 10 11 54 AM (1)](https://github.com/user-attachments/assets/6cf0ebef-173b-400a-8ce8-8a3950da1d34)
![WhatsApp Image 2025-06-25 at 10 11 54 AM](https://github.com/user-attachments/assets/fa4e6955-7bbe-4643-b4ca-9b9f76b447a7)



🌐 HOW ROUTING WORKS IN THIS PROJECT:

Nginx reverse proxy handles all incoming requests and decides which backend app (Go or Python) should respond, based on the URL path.


🔁 Step-by-step Flow:

1. User sends a request to:

http://localhost/service1/hello → Go app

http://localhost/service2/hello → Python app



2. Nginx receives the request on port 80.


3. Inside nginx.conf, we define path-based routing:

location /service1/ {
    rewrite ^/service1(/.*)$ $1 break;
    proxy_pass http://service1:8001;
}

location /service2/ {
    rewrite ^/service2(/.*)$ $1 break;
    proxy_pass http://service2:8002;
}


4. Explanation:

location /service1/: matches any request that starts with /service1/.

rewrite ^/service1(/.*)$ $1 break;: removes /service1 from the path (so Go app gets just /hello).

proxy_pass http://service1:8001;: forwards the request to the Go app container named service1 on port 8001.

Same logic applies for /service2 and the Python app.



5. The apps are running in bridge network, so Nginx can reach service1 and service2 using the service/container name.


6. Response from app → goes back through Nginx → to the user.
