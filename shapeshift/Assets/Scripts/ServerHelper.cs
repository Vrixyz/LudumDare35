using UnityEngine;
using System.Collections;

using System;
using System.Text;
using System.Net;
using System.Net.Sockets;

public class ServerHelper {
    ServerSender serverSender;

    public string ip = "127.0.0.1";  // define in init
    public int port = 10003;  // define in init
    
    public void init()
    {
        serverSender = new ServerSender();


        serverSender.init(ip, port);

        // testing via console
        // sendObj.inputFromConsole();
    }
    public void sendString(string message)
    {
        serverSender.sendString(message + "\n");
    }

    [Serializable]
    private class MoveMessage
    {
        public float XSpeed;
        public float YSpeed;
    }

    /// <summary>
    /// high level helper for moving
    /// </summary>
    public void move(float x, float y)
    {
        MoveMessage moveMessage = new MoveMessage();
        moveMessage.XSpeed = x;
        moveMessage.YSpeed = y;
        string json = JsonUtility.ToJson(moveMessage);
        serverSender.sendString(json);
    }

}
