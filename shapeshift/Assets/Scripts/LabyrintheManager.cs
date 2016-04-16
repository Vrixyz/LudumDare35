using UnityEngine;
using System.Collections;

public class LabyrintheManager : MonoBehaviour
{
    public static CaseLab[,] lab;
    public static float CASE_SIZE = 1f;

    private GameObject player;

    private int countX;
    private int countY;

    private string[] entry = new string[] { "0000000000",
                            "0000110110",
                            "1111100110",
                            "0000101111",
                            "0010001012",
                            "1101101010",
                            "0001000000",
                            "0001011111",
                            "0111000000",
                            "0000011111" };

    void Start()
    {
        initLab();
    }

    private void initLab()
    {
        countX = entry.Length;
        countY = entry[0].Length;
        lab = new CaseLab[countX, countY];
        for (int i = 0; i < countX; i++)
        {
            string line = entry[i];
            for (int j = 0; j < countY; j++)
            {
                int curr_value = 0;
                int.TryParse(line.Substring(j, 1), out curr_value);
                lab[i, j] = new CaseLab(curr_value);
                if(curr_value == 0)
                {
                    lab[i, j].element = createPathAtPosition(i, j);
                }
                else if (curr_value == 1)
                {
                    lab[i, j].element = createWallAtPosition(i, j);
                }
                else if (curr_value == 2)
                {
                    lab[i, j].element = createPathAtPosition(i, j);
                    player = createPlayerAtPosition(i, j);
                    player.AddComponent<Rigidbody>();
                    player.AddComponent<PlayerScript>();
                }
            }
        }
    }
    
    private GameObject createPathAtPosition(float x, float z)
    {
        GameObject go = GameObject.CreatePrimitive(PrimitiveType.Plane);
        go.transform.SetParent(transform);
        go.transform.position = new Vector3(x + CASE_SIZE / 2, 0, z + CASE_SIZE / 2);
        //go.transform.position = new Vector3(x, 0, z);
        go.transform.localScale = Vector3.one * CASE_SIZE / 10;
        return go;
    }

    private GameObject createWallAtPosition(float x, float z)
    {
        GameObject go = GameObject.CreatePrimitive(PrimitiveType.Cube);
        go.transform.SetParent(transform);
        go.transform.localScale = Vector3.one * CASE_SIZE;
        go.transform.position = new Vector3(x + CASE_SIZE / 2, go.transform.localScale.y / 2, z + CASE_SIZE / 2);
        //go.transform.position = new Vector3(x, 0, z);
        return go;
    }

    private GameObject createPlayerAtPosition(float x, float z)
    {
        GameObject go = GameObject.CreatePrimitive(PrimitiveType.Sphere);
        go.transform.SetParent(transform);
        go.transform.localScale = Vector3.one * CASE_SIZE;
        go.transform.position = new Vector3(x + CASE_SIZE / 2, go.transform.localScale.y / 2, z + CASE_SIZE / 2);
        //go.transform.position = new Vector3(x, 0, z);
        return go;
    }

    void OnRenderObject()
    {
        MaterialsManager.getLineMaterial().SetPass(0);

        GL.PushMatrix();
        GL.MultMatrix(Matrix4x4.identity);
        
        GL.Begin(GL.LINES);
        for (int i = 0; i <= Mathf.Max(countX, countY); i++)
        {
            if(i <= countX)
            {
                GL.Vertex3(i, 0, 0);
                GL.Vertex3(i, 0, countY);
            }
            if (i <= countY)
            {
                GL.Vertex3(0, 0, i);
                GL.Vertex3(countX, 0, i);
            }
        }
        GL.End();
        GL.PopMatrix();
    }
}

public class CaseLab
{
    public GameObject element { get; set; }
    public int value { get; set; }

    public CaseLab(GameObject element, int value)
    {
        this.element = element;
        this.value = value;
    }

    public CaseLab(int value)
    {
        this.value = value;
    }
}
