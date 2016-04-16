using UnityEngine;
using System.Collections;

public class ServerHelperGUI : MonoBehaviour {
    ServerHelper serverHelper = new ServerHelper();
    ServerListener serverListener = new ServerListener();

    // gui
    string strMessage = "{\"Xspeed\":1}";

    // Use this for initialization
    void Start () {
        init();
    }

    // Use this for initialization
    void init()
    {
        serverHelper.init();
        serverListener.init();
    }

    // OnGUI
    void OnGUI()
    {
        /// sender part

        Rect rectObj = new Rect(40, 380, 200, 400);
        GUIStyle style = new GUIStyle();
        style.alignment = TextAnchor.UpperLeft;
        GUI.Box(rectObj, "# UDPSend-Data\n" + serverHelper.ip + ":" + serverHelper.port + " #\n"
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
        GUI.Box(rectObj2, "# UDPReceive\n" + serverHelper.ip + ":" + serverListener.port + " #\n"
                    + "shell> nc -u " + serverListener.ip + ":" + serverListener.port + " \n"
                    + "\nLast Packet: \n" + serverListener.lastReceivedUDPPacket
                    + "\n\nAll Messages: \n" + serverListener.allReceivedUDPPackets
                , style);
    }
}
