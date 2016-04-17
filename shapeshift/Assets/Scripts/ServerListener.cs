using UnityEngine;
using System.Collections;

using System;
using System.Text;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Globalization;

public class ServerListener {

    public delegate void OnReceive(byte[] data);

    public OnReceive onReceiveDelegate;

    // receiving Thread
    Thread receiveThread;

    // udpclient object
    public UdpClient client;
    
    string ip;
    int port; // define > init

    // init
    public void init(string p_ip, int p_port)
    {
        ip = p_ip;
        port = p_port;
        client = new UdpClient(p_port);
        receiveThread = new Thread(
            new ThreadStart(ReceiveData));
        receiveThread.IsBackground = true;
        receiveThread.Start();
    }
    int abort = 0;
    // receive thread
    private void ReceiveData()
    {
        while (true)
        {

            try
            {
                IPEndPoint serverIP = new IPEndPoint(IPAddress.Any, 0);
                byte[] data = client.Receive(ref serverIP);
                onReceiveDelegate(data);
            }
            catch (Exception err)
            {
                Debug.Log(err);
            }
        }
    }

    public void stop()
    {
        // end of application
        if (receiveThread != null)
        {
            receiveThread.Abort();
            if (client != null)
            {
                client.Close();
            }
        }
    }
}
