#include <iostream>
#include <cstring>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>

int main() {
    int server_socket, client_socket;
    struct sockaddr_in server_addr, client_addr;

    // Create a socket
    server_socket = socket(AF_INET, SOCK_STREAM, 0);

    // Configure the server address structure
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(12345);  // Choose a port
    server_addr.sin_addr.s_addr = INADDR_ANY;

    // Bind the socket to the server address
    bind(server_socket, (struct sockaddr*)&server_addr, sizeof(server_addr));

    // Listen for incoming connections
    listen(server_socket, 5);  // Number of queued connections

    std::cout << "Server listening on port 12345..." << std::endl;

    // Accept an incoming connection
    socklen_t client_addr_len = sizeof(client_addr);
    client_socket = accept(server_socket, (struct sockaddr*)&client_addr, &client_addr_len);

    std::cout << "Accepted connection from " << inet_ntoa(client_addr.sin_addr) << std::endl;

    // Send a welcome message to the client
    const char* welcome_message = "Welcome to the server!\n";
    send(client_socket, welcome_message, strlen(welcome_message), 0);

    // Close the sockets
    close(client_socket);
    close(server_socket);

    return 0;
}
