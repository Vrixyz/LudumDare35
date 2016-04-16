using UnityEngine;
using System.Collections;
using System.Globalization;
using System.Text;
using System;


public class ServerHelperGUI : MonoBehaviour {
    ServerHelper serverHelper = new ServerHelper();

    // gui
    string strMessage = "{\"Xspeed\":1}";

    // infos
    public string lastReceivedUDPPacket = "";
    public string allReceivedUDPPackets = ""; // clean up this from time to time!

    // Use this for initialization
    void Start () {
        init();
    }

    // Use this for initialization
    void init()
    {
        serverHelper.onReceivePlayersDelegate = (playersMessage) =>
        {
            string players = JsonUtility.ToJson(playersMessage);
            print(">> " + players);
            // latest UDPpacket
            lastReceivedUDPPacket = players;

            // ....
            allReceivedUDPPackets = allReceivedUDPPackets + " ; " + players;
        };
        serverHelper.init();
    }

    // OnGUI
    void OnGUI()
    {
        /// sender part

        Rect rectObj = new Rect(40, 380, 200, 400);
        GUIStyle style = new GUIStyle();
        style.alignment = TextAnchor.UpperLeft;
        GUI.Box(rectObj, "# UDPSend-Data\n" + serverHelper.ip + ":" + serverHelper.sendPort + " #\n"
                , style);
        // ------------------------
        // send it
        // ------------------------
        strMessage = GUI.TextField(new Rect(40, 420, 140, 20), strMessage);
        if (GUI.Button(new Rect(190, 420, 40, 20), "send"))
        {
            serverHelper.sendString(strMessage);
        }


        /// listener part
        
        Rect rectObj2 = new Rect(40, 10, 200, 400);
        style.alignment = TextAnchor.UpperLeft;
        GUI.Box(rectObj2, "# UDPReceive\n" + serverHelper.ip + ":" + serverHelper.listenPort + " #\n"
                    + "shell> nc -u " + serverHelper.ip + ":" + serverHelper.listenPort + " \n"
                    + "\nLast Packet: \n" + lastReceivedUDPPacket
                    + "\n\nAll Messages: \n" + allReceivedUDPPackets
                , style);
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
        serverHelper.stop();
    }
}
