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
    UdpClient client;
    
    string ip;
    int port; // define > init

    // init
    public void init(string p_ip, int p_port)
    {
        ip = p_ip;
        port = p_port;

        receiveThread = new Thread(
            new ThreadStart(ReceiveData));
        receiveThread.IsBackground = true;
        receiveThread.Start();

    }

    // receive thread
    private void ReceiveData()
    {

        client = new UdpClient(port);
        while (true)
        {

            try
            {
                IPEndPoint serverIP = new IPEndPoint(IPAddress.Parse(ip), port);
                byte[] data = client.Receive(ref serverIP);
                onReceiveDelegate(data);
            }
            catch (Exception err)
            {
                //TODO: handle
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
