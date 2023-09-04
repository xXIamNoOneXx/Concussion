import tkinter as tk
import socket
import logging
from cryptography.fernet import Fernet

# Configure the logging module
logging.basicConfig(filename='c2_client.log', level=logging.INFO, format='%(asctime)s - %(levelname)s: %(message)s')

class C2Client:
    def __init__(self, root):
        self.root = root
        root.title("Concussion C2 Client")

        self.server_host = 'YOUR_SERVER_IP'
        self.server_port = 12345  # Match the server's port
        self.client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            # Attempt to connect to the C2 server
            self.client_socket.connect((self.server_host, self.server_port))
        except ConnectionRefusedError:
            logging.error("Connection refused. Ensure the C2 server is running.")
            self.root.destroy()
        except Exception as e:
            logging.error(f"Error: {e}")
            self.root.destroy()

        # Load the shared symmetric key (you should securely distribute this key from the server)
        self.symmetric_key = b'YOUR_SHARED_SYMMETRIC_KEY'
        self.cipher_suite = Fernet(self.symmetric_key)

        self.setup_gui()

    def setup_gui(self):
        # Create and configure your GUI elements here
        # ...

    def send_command(self):
        # Implement command sending logic here
        # ...

    def show_history(self):
        # Implement command history retrieval logic here
        # ...

    def close_connection(self):
        # Log connection closure when the GUI is closed
        logging.info("Connection to the C2 server closed.")
        self.client_socket.close()

if __name__ == "__main__":
    root = tk.Tk()
    client = C2Client(root)

    # Configure the close button to log connection closure
    root.protocol("WM_DELETE_WINDOW", client.close_connection)

    root.mainloop()
