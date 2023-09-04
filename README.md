- 👋 Hi, I’m @xXIamNoOneXx
- 👀 I’m interested in ...everything computer
- 🌱 I’m currently learning ...all that I can
- 💞️ I’m looking to collaborate on ...anything...well almost anything.
  
- 📫 How to reach me ... nullbitt5000@gmail.com

This is currently a work in progress. I am currently learning about command and control servers and I always feel like I learn the most when I do it so I started making one. Any feedback or suggestions are welcome. Thank YOU
P.S. This is for educational purposes only. What you do with my code from here is your responsibility.

Description of "Concussion" C2 System:

"Concussion" is a simple command and control (C2) system consisting of a server (c2_server.py), a client with a GUI (c2_client.py), and an optional simple C++ server (simple_server.cpp) for demonstration purposes. Here's what it can do:

    Server (c2_server.py):
        Listens for incoming connections.
        Accepts encrypted commands from the client.
        Executes commands on the server's machine.
        Sends encrypted command output back to the client.
        Supports command history and an "exit" command to close the connection.

    Client with GUI (c2_client.py):
        Provides a user-friendly interface for sending commands to the server.
        Sends encrypted commands to the server.
        Displays the encrypted output received from the server.
        Allows viewing command history.
        Provides an "exit" button to close the connection gracefully.

    Optional Simple C++ Server (simple_server.cpp):
        A demonstration server that listens on port 12345.
        Accepts connections and sends a welcome message.

Installation Instructions:

Server (c2_server.py):

    Place c2_server.py on your server machine.
    Install the required Python packages by running pip install cryptography.
    Configure the server's host and port in the script.
    Generate a shared symmetric key and replace 'YOUR_SHARED_SYMMETRIC_KEY' with it.
    Run the server script with python c2_server.py.

Client with GUI (c2_client.py):

    Place c2_client.py on your client machine.
    Install the required Python packages by running pip install cryptography tkinter.
    Configure the server's host and port, and replace 'YOUR_SERVER_IP' and 12345 with the server's details.
    Replace 'YOUR_SHARED_SYMMETRIC_KEY' with the same symmetric key used on the server.
    Run the client script with python c2_client.py.

Optional Simple C++ Server (simple_server.cpp):

    Compile simple_server.cpp using a C++ compiler.
    Place the compiled executable on a machine where you want to run the simple server.
    Run the server executable.

Ensure that the server and client machines can communicate over the network and that you have proper authorization for any activities involving system manipulation or penetration testing.
