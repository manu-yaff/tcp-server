package main

import (
	"net"
	"strings"
	"tcp-server/channel"
	c "tcp-server/client"
	s "tcp-server/server"
	"testing"
)

var server = &s.Server{
	Clients:  make(map[net.Conn]*c.Client),
	Channels: make(map[string]*channel.Channel),
}

var conn net.Conn

func initServer() {
	server.StartServer("2222")
	server.ListenForConnections()
}

func TestStartServer(t *testing.T) {
	go initServer()
	c.ConnectToServer("localhost", "2222")
}

func TestAddClientToLobby(t *testing.T) {
	newClient := server.AddClientToLoby(&conn)

	if newClient.Name != "Anonymus" {
		t.Errorf("%s is different than %s", newClient.Name, "Anonymus")
	}

	if newClient.Conn != conn {
		t.Errorf("%s is different than %s", newClient.Conn, conn)
	}

	if len(server.Clients) != 1 {
		t.Errorf("%d is different than %d", len(server.Clients), 1)
	}
}

func TestSetClientName(t *testing.T) {
	server.AddClientToLoby(&conn)
	server.SetClientName("Jon", conn)
	if server.Clients[conn].Name != "Jon" {
		t.Errorf("%s is different than %s", "Jon", server.Clients[conn].Name)
	}
}

func TestGetChanne(t *testing.T) {
	server.Channels["general"] = &channel.Channel{
		Name: "general",
	}
	server.Channels["dev"] = &channel.Channel{
		Name: "dev",
	}

	expectedChannels := 2
	temp := server.GetChannels()
	actualChannels := strings.Split(temp, ",")

	if expectedChannels != len(actualChannels) {
		t.Errorf("got %s, expected %d", actualChannels, expectedChannels)
	}

}
