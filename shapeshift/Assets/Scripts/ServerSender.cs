﻿/*
 
    -----------------------
    UDP-Send
    -----------------------
    // [url]http://msdn.microsoft.com/de-de/library/bb979228.aspx#ID0E3BAC[/url]
   
    // > gesendetes unter
    // 127.0.0.1 : 8050 empfangen
   
    // nc -lu 127.0.0.1 8050
 
        // todo: shutdown thread at the end
*/
using UnityEngine;
 
using System;
using System.Text;
using System.Net;
using System.Net.Sockets;
 
public class ServerSender
{
    private static int localPort;
   
    // "connection" things
    IPEndPoint remoteEndPoint;
    public UdpClient client;

    // init
    public void init(UdpClient p_client, string ip, int port)
    {
        remoteEndPoint = new IPEndPoint(IPAddress.Parse(ip), port);
        client = p_client;
        client.Connect(remoteEndPoint);
    }
 
    // sendData
    public void sendString(string message)
    {
        try
        {
                //if (message != "")
                //{
                    byte[] data = Encoding.UTF8.GetBytes(message);
                    client.Send(data, data.Length);
                //}
        }
        catch (Exception err)
        {
            Debug.Log(err);
            // TODO: handle this
        }
    }
}