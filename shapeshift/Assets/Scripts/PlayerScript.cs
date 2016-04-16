using UnityEngine;
using System.Collections;

public class PlayerScript : MonoBehaviour
{
    Rigidbody rb;
    private float speed = 5;
    void Start()
    {
        rb = GetComponent<Rigidbody>();
        rb.mass = 0;
        rb.useGravity = false;
        rb.constraints = RigidbodyConstraints.FreezePositionY | RigidbodyConstraints.FreezeRotation;
        
    }
    Vector3 move = Vector3.zero;
    void FixedUpdate ()
    {
        float inputX = Input.GetAxisRaw("Horizontal");
        float inputY = Input.GetAxisRaw("Vertical");

        if (inputX != 0 || inputY != 0)
        {
            move = Vector3.zero;

            move.x += inputX * speed * Time.deltaTime;
            move.z += inputY * speed * Time.deltaTime;

            rb.MovePosition(transform.position + move);
            //transform.position += move;
        }
        else
        {
            if (move.magnitude > 0.09f)
            {
                move -= move * Time.deltaTime;
                transform.position += move;
            }
            else
            {
                move = Vector3.zero;
            }
        }
        rb.velocity = Vector3.zero;
    }
}
