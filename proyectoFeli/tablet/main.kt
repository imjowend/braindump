import android.media.MediaPlayer
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.view.View
import android.widget.Button
import android.widget.EditText

class MainActivity : AppCompatActivity() {

    private lateinit var input: EditText
    private lateinit var mediaPlayer: MediaPlayer

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        // Obtiene referencias a los elementos de la interfaz
        input = findViewById(R.id.input)
        mediaPlayer = MediaPlayer.create(this, R.raw.ambient_42)

        // Asigna los listeners a los botones
        setButtonListener(R.id.button_1, "1")
        setButtonListener(R.id.button_2, "2")
        setButtonListener(R.id.button_3, "3")
        setButtonListener(R.id.button_4, "4")
        setButtonListener(R.id.button_5, "5")
        setButtonListener(R.id.button_6, "6")
        setButtonListener(R.id.button_7, "7")
        setButtonListener(R.id.button_8, "8")
        setButtonListener(R.id.button_9, "9")
        setButtonListener(R.id.button_0, "0")
        setButtonListener(R.id.button_star, "*")
        setButtonListener(R.id.button_hash, "#")
        setButtonListener(R.id.button_clear, "clear")
    }

    // Asigna el listener a un botón
    private fun setButtonListener(id: Int, value: String) {
        val button = findViewById<Button>(id)
        button.setOnClickListener { view: View? ->
            when (value) {
                "clear" -> clearDigit()
                else -> insert(value)
            }
            playSound()
        }
    }

    // Función para insertar un dígito en el input
    private fun insert(digit: String) {
        input.append(digit)
    }

    // Función para borrar el último dígito del input
    private fun clearDigit() {
        val text = input.text.toString()
        if (text.isNotEmpty()) {
            input.setText(text.substring(0, text.length - 1))
        }
    }

    // Función para reproducir un sonido
    private fun playSound() {
        mediaPlayer.seekTo(0)
        mediaPlayer.start()
    }
}
