using UnityEngine;
using System.Collections;

using System;
using System.Text;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Globalization;

public class ServerListener : MonoBehaviour {

    [Serializable]
    public class Player
    {
        public string name;
        public Vector2 position;
    }
	[Serializable]
    public class ObjectMessage
    {
		public string time;
        public Player[] players;
    }

    // receiving Thread
    Thread receiveThread;

    // udpclient object
    UdpClient client;

    // public
    public string ip = "127.0.0.1";
    public int port = 10002; // define > init

    // infos
    public string lastReceivedUDPPacket = "";
    public string allReceivedUDPPackets = ""; // clean up this from time to time!

    // start from unity3d
    public void Start()
    {

        init();
    }

    // init
    public void init()
    {
        // Endpunkt definieren, von dem die Nachrichten gesendet werden.
        print("UDPSend.init()");

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
                // Bytes empfangen.
                IPEndPoint anyIP = new IPEndPoint(IPAddress.Parse("ip"), port);
                byte[] data = client.Receive(ref anyIP);

                // Bytes mit der UTF8-Kodierung in das Textformat kodieren.
                string text = Encoding.UTF8.GetString(data);

                // Den abgerufenen Text anzeigen.
                print(">> " + text);
                Player player = JsonUtility.FromJson<Player>(text);
                print(">> " + player.name);
                // latest UDPpacket
                lastReceivedUDPPacket = text;

                // ....
                allReceivedUDPPackets = allReceivedUDPPackets + text;

            }
            catch (Exception err)
            {
                print(err.ToString());
            }
        }
    }

    // getLatestUDPPacket
    // cleans up the rest
    public string getLatestUDPPacket()
    {
        allReceivedUDPPackets = "";
        return lastReceivedUDPPacket;
    }
	
	public void OnApplicationQuit()
     {
        // end of application
        if (receiveThread != null)
        { 
			receiveThread.Abort();
			if (client != null) {
				client.Close();
			}
        }
  
        print("Stop"); 
     } 
}
