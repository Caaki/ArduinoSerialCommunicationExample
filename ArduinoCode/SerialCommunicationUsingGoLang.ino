/*
  Blink

  Turns an LED on for one second, then off for one second, repeatedly.

  Most Arduinos have an on-board LED you can control. On the UNO, MEGA and ZERO
  it is attached to digital pin 13, on MKR1000 on pin 6. LED_BUILTIN is set to
  the correct LED pin independent of which board is used.
  If you want to know what pin the on-board LED is connected to on your Arduino
  model, check the Technical Specs of your board at:
  https://www.arduino.cc/en/Main/Products

  modified 8 May 2014
  by Scott Fitzgerald
  modified 2 Sep 2016
  by Arturo Guadalupi
  modified 8 Sep 2016
  by Colby Newman

  This example code is in the public domain.

  https://www.arduino.cc/en/Tutorial/BuiltInExamples/Blink
*/

String command;
// the setup function runs once when you press reset or power the board
void setup() {
  Serial.begin(9600);
  // initialize digital pin LED_BUILTIN as an output.
  pinMode(LED_BUILTIN, OUTPUT);
  delay(2000);
  Serial.println("ON/OFF");

}

// the loop function runs over and over again forever
void loop() {

  if (Serial.available()){
    command = Serial.readStringUntil('\n');
    command.trim();
  
    if (command.equals("on")){
      digitalWrite(LED_BUILTIN, HIGH); 
      Serial.write("LED IS ON"); // turn the LED on (HIGH is the voltage level)
    }else if(command.equals("off")){
      digitalWrite(LED_BUILTIN, LOW);
        Serial.write("LED IS OFF");   // turn the LED off by making the voltage LOW
    }
    else{
      Serial.write("Bad command");
      digitalWrite(LED_BUILTIN, HIGH);
      delay(200);                      // wait for a second
      digitalWrite(LED_BUILTIN, LOW);
      delay(200);
      digitalWrite(LED_BUILTIN, HIGH);
      delay(200);
      digitalWrite(LED_BUILTIN, LOW);
      delay(200);
    }
  }

}
