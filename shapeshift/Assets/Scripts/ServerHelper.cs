using UnityEngine;
using System.Collections;

using System;
using System.Text;
using System.Net;
using System.Net.Sockets;

public class ServerHelper {
    ServerSender serverSender = new ServerSender();
    ServerListener serverListener = new ServerListener();

    public string ip = "163.172.27.180";  // define in init
    public int sendPort = 10003;  // define in init
    public int listenPort = 10002;  // define in init

    public void init()
    {
        serverListener.onReceiveDelegate = (data) =>
        {
            string text = Encoding.UTF8.GetString(data);

            PlayersMessage playersMessage = JsonUtility.FromJson<PlayersMessage>(text);
            onReceivePlayersDelegate(playersMessage);
        };
        serverListener.init(ip, listenPort);

        serverSender.init(ip, sendPort);
    }
    public void sendString(string message)
    {
        serverSender.sendString(message + "\n");
    }


    /// <summary>
    /// Vector2 from my Go implementation
    /// </summary>
    [Serializable]
    public class Vector2Go
    {
        public float X;
        public float Y;
    }
    /// <summary>
    /// Struct you receive to display a player
    /// </summary>
    [Serializable]
    public class Player
    {
        public string Id;
        public string Name;
        public Vector2Go Position;
    }
    /// <summary>
    /// Struct you receive to display players
    /// </summary>
    [Serializable]
    public class PlayersMessage
    {
        public string time;
        public Player[] players;
    }

    public delegate void OnReceivePlayers(PlayersMessage data);
    public OnReceivePlayers onReceivePlayersDelegate;



    /// <summary>
    /// struct to send for a move action
    /// </summary>
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

    public void stop()
    {
        serverListener.stop();
    }
}
