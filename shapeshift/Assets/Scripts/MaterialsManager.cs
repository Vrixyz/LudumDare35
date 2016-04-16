using UnityEngine;
using System.Collections;

public class MaterialsManager : MonoBehaviour
{
    private static Material lineMaterial;
    public static void createLineMaterial()
    {
        var shader = Shader.Find("Hidden/Internal-Colored");
        lineMaterial = new Material(shader);
        lineMaterial.hideFlags = HideFlags.HideAndDontSave;
        lineMaterial.SetInt("_SrcBlend", (int)UnityEngine.Rendering.BlendMode.SrcAlpha);
        lineMaterial.SetInt("_DstBlend", (int)UnityEngine.Rendering.BlendMode.OneMinusSrcAlpha);
        lineMaterial.SetInt("_Cull", (int)UnityEngine.Rendering.CullMode.Off);
        lineMaterial.SetInt("_ZWrite", 0);
    }

    public static Material getLineMaterial()
    {
        if (!lineMaterial)
        {
            createLineMaterial();
        }
        return lineMaterial;
    }
}
