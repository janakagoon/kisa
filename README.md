# Dexterous Actuators

## Demonstration Video (Click to Play) 

[![Demonstration Video of Robotic Actuator](https://img.youtube.com/vi/waDG-JWzCfU/0.jpg)](https://www.youtube.com/watch?v=waDG-JWzCfU)

## Fly-Through Videos (Click to Play)

| | | |
|-|-|-|
|[![Demonstration Video of Robotic Actuator](https://img.youtube.com/vi/6SOCoPpx5d4/0.jpg)](https://www.youtube.com/watch?v=6SOCoPpx5d4)|[![Demonstration Video of Robotic Actuator](https://img.youtube.com/vi/lXHCQ28MPKY/0.jpg)](https://www.youtube.com/watch?v=lXHCQ28MPKY)|[![Demonstration Video of Robotic Actuator](https://img.youtube.com/vi/2JxnCmN5SNg/0.jpg)](https://www.youtube.com/watch?v=2JxnCmN5SNg)|

## Appendix A: List of Supported Grips

| | | |
|-|-|-|
<img src="https://github.com/janakagoon/kisa/blob/main/misc/images/robot-cat-tree-small.png?raw=true" height="150"> |<img src="https://github.com/janakagoon/kisa/blob/main/misc/images/robot-supermarket-small.png?raw=true" height="150"> |<img src="https://github.com/janakagoon/kisa/blob/main/misc/images/robot-kitchen-small.png?raw=true" height="150"> |

Cutkosky and Wright (1986b) studied various hand grasps used by technicians in a machine shop. Based on their observations they generated a taxonomy for human grips. They identified 16 grasps used in such a work environment.

We have listed these below and noted whether or not they can be performed by the robotic actuator (RA). 

Terms:
* Prehensile: Grabbing with fingers.
* Non-prehensile: Involves no grabbing.
* Prismatic: Grabbing like a cylinder.
* Circular: Grabbing like a ball.
* Power: Grip formed by the inner surfaces of the fingers and palm.
* Precision: the grip is formed by the fingertips.

<small>The illustrations of the graphs are taken from "Human Grasp Choices and Robotic Grasp Analysis" by Cutkosky and Howe. Copyright belongs to the original owner M. J. Dowling and Robotics Institute, Carnegie-Mellon University</small>

| Grasp | Illustration | Type | Can it be performed by the Actuator? | RA Image |
| :--- | :--- | :--- | :--- | :--- |
| Heavy Wrap - Large Diameter |![](https://raw.githubusercontent.com/janakagoon/kisa/main/misc/grips/illustrations/1%20-%20Heavy%20Wrap%20-%20Large%20Diameter.png?raw=true)| Power, Prehensile, Prismatic | Yes | ![](https://raw.githubusercontent.com/janakagoon/kisa/main/misc/grips/images/1%20-%20Heavy%20Wrap%20-%20Large%20Diameter.png)|
| Heavy Wrap - Small Diameter |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/2%20-%20Heavy%20Wrap%20-%20Small%20Diameter.png?raw=true)| Power, Prehensile, Prismatic | No | |
| Medium Wrap |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/3%20-%20Medium%20Wrap.png?raw=true)| Power, Prehensile, Prismatic | No | |
| Abducted Thumb |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/4%20-%20Abducted%20Thumb.png?raw=true)| Power, Prehensile, Prismatic | No | |
| Light Tool |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/5%20-%20Light%20Tool.png?raw=true)| Power, Prehensile, Prismatic | Yes | ![](https://raw.githubusercontent.com/janakagoon/kisa/main/misc/grips/images/5%20-%20Light%20Tool.png) |
| Thumb-4 Finger |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/6%20-%20Thumb-4%20Fingers.png?raw=true)| Precision, Prismatic | No | |
| Thumb-3 Finger |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/7%20-%20Thumb-3%20Fingers.png?raw=true)| Precision, Prismatic | No | |
| Thumb-2 Finger |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/8%20-%20Thumb-2%20Fingers.png?raw=true)| Precision, Prismatic| Yes | ![](https://github.com/janakagoon/kisa/blob/main/misc/grips/images/8%20-%20Thumb-2%20Fingers.png?raw=true) |
| Thumb-Index Finger |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/9%20-%20Thumb-Index%20Finger.png?raw=true)| Precision, Prismatic | Yes | ![](https://github.com/janakagoon/kisa/blob/main/misc/grips/images/9%20-%20Thumb-Index%20Finger.png?raw=true) |
| Power-Disk |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/10%20-%20Power%20Disk.png?raw=true)| Power, Circular | Yes | ![](https://github.com/janakagoon/kisa/blob/main/misc/grips/images/10%20-%20Power%20Disk.png?raw=true) |
| Power Sphere |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/11%20-%20Power%20Sphere%20.png?raw=true)| Power, Circular | Yes | Same as Power-Disk, except with more flex on the end joint. |
| Precision Disk |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/12%20-%20Precision%20Disk.png?raw=true)| Precision, Circular | Yes | ![](https://github.com/janakagoon/kisa/blob/main/misc/grips/images/12%20-%20Precision%20Disk.png?raw=true) |
| Precision Sphere |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/13%20-%20Precision%20Sphere.png?raw=true)| Precision, Circular | Yes | Similar to Precision disk, but joints are more closed. |
| Tripod |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/14%20-%20Tripod.png?raw=true)| Precision, Circular | Yes | ![](https://github.com/janakagoon/kisa/blob/main/misc/grips/images/14%20-%20Tripod.png?raw=true) |
| Platform |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/15%20-%20Platform.png?raw=true)| Power, Non-Prehensile | Yes |  ![](https://github.com/janakagoon/kisa/blob/main/misc/grips/images/15%20-%20Platform.png?raw=true)|
| Lateral Pinch |![](https://github.com/janakagoon/kisa/blob/main/misc/grips/illustrations/16%20-%20Lateral%20Pinch.png?raw=true)| Power, Prehensile, Prismatic | No | |


## Appendix B: How to Build a 32-Motor Motor Controller

For the robotic actuator to perform grips, we need to power the motors in the arm in a coordinated manner. To enable this we built a programmable motor controller which can control up to 32 electric motors.

We believe this device and the software libraries written to command it, to be generally useful, and thus have provided the design for the controller and the software here for free.

![](https://github.com/janakagoon/kisa/blob/main/misc/images/controller-and-hand.png?raw=true)

![](https://github.com/janakagoon/kisa/blob/main/misc/images/controller-close-up.png?raw=true)

To build the controller you will need:

| | | |
|-|-|-|
| x1 | Raspberry Pi | ![](https://github.com/janakagoon/kisa/blob/main/misc/images/raspberry-pi-top-view.png?raw=true) |
| x5 | Eight relays 8-Layer Stackable HAT for Raspberry Pi from Sequant Microsystems. (https://sequentmicrosystems.com/products/raspberry-pi-relays-stackable-card) | ![](https://github.com/janakagoon/kisa/blob/main/misc/images/8-relay-board.png?raw=true) |
| 10 feet of wires. Ideally of 4 different colors. | Used to connect the relays in to H-Bridges, that control the motors. | ![](https://github.com/janakagoon/kisa/blob/main/misc/images/h-bridges-closeup.png?raw=true) |


The relay boards communicate with the raspberry pi via the I2C protocol. Each board has  3 jumper switches used to assign each board a unique device identifier. The boards can be assigned sequential IDs by setting their respective jumpers as shown:

![](https://github.com/janakagoon/kisa/blob/main/misc/images/relay-board-jumpers.png?raw=true)

The relay boards are then stacked on top of the Raspberry Pi, connecting to the Pi's GPIO connector pins.


The Raspberry PI OS comes with an I2C driver, however it is disabled by default and must be turned on, in the boot settings. See here for details: https://learn.adafruit.com/adafruits-raspberry-pi-lesson-4-gpio-setup/configuring-i2c

Once enabled, this device file will be created in the file system: `/dev/i2c-1`

An application running on the Raspberry Pi can get the state of a relay and turn a relay on/off by reading and writing byte values to and from an offset into this file. The offset is calculated using the board id - a number in the range of 0 and 5 - and the relay id, of each board - a number in the range of 1 and 8. The details on how to do this offset calculation is found in the github project shared below.

Pairs of relays are then wired together to make H-bridges which are then directly attached to the arm motors and an appropriate power supply.

We used this power supply: BMOUO 12V 30A DC Universal Regulated Switching Power Supply (https://www.amazon.com/gp/product/B01EWG6YT8).

Each relay board has 8 relays, so you can make 4 H-bridges on each board.

This power supply provides -12V and +12V connectors, so together they can power motors requiring 24V DC power source. 

The schematic of the H-Bridge:

![](https://github.com/janakagoon/kisa/blob/main/misc/images/relay-schema.png?raw=true)

You will need to wire together pairs of relays on each board as shown, according to the above schema:

![](https://github.com/janakagoon/kisa/blob/main/misc/images/detail-relay-to-h-bridge.png?raw=true)

* Relay 1 and relay 2's +12 volt connector - the red wires - are connected together and then plugged in to the power supply's +12 connector
* Relay 1 and relay 2's -12 volt connector - the blue wires - are connected together and plugged in to the power supply's -12 V connector
* The relay 1's COM connector - the white wire - and relay 2's COM connector - the black wire - are connected to the motor's terminals.

WARNING! WARNING! WARNING!

Plugging in the COM wires in to the power supply will destroy the relay board!

So make sure you don't do that accidentally.

WARNING! WARNING! WARNING!

### Control Software

To set and get a relay, start from this file:


https://github.com/janakagoon/kisa/blob/main/lib/relay.go


The data types for defining H-bridges and motors as high level objects, see here:


https://github.com/janakagoon/kisa/blob/main/lib/motion.go


This was specific to the robotic end actuator, but you could probably copy and past parts of it for most other applications. 


## Appendix C: Solution To The QR-Code Based Optical Configuration Detection Problem.

Optical configuration detection consists of three steps:

1. Detecting all the observable QR codes in a video frame.
2. Calculating the position of each QR code in the World coordinate system.
3. Solving for the joint flex angles of the hand, using position of the QR codes obtained in step 2.

Step 1 There are cloud services that use Machine Learning models to accomplish rapid and accurate detection of QR codes in a video feed. The output of these will be the bounding box of each QR code detected in the video frame.

![](https://github.com/janakagoon/kisa/blob/main/misc/images/QR-code-annotation.png?raw=true)

Step 2 is achieved by solving the Camera Equation below, using the results of Step 1.

![](https://github.com/janakagoon/kisa/blob/main/misc/images/camera-equation.png?raw=true)

where:
* `(u, v)` is the location of the QR code in the video frame.
* `f_x`, `f_y` are the focal lengths of the camera.
* `u_0` and `v_0` and camera distortion parameters which can be obtained by the manufacturer.
* `s`  is a scaling factor in the image, which would be known, form the camera zoom setting.
* The `[r_ij]` matrix is the camera's rotation matrix and `(t_i)` vector is the camera's translation vector from the World coordinate system's origin. These can be measured directly. If the camera is fixed on the workbench, these will be static.
* `(x,y,z)` are the coordinates of the QR code in the world coordinate system. Solving this equation will yield a line in the World coordinate system. The line will pass through the camera.

The location of the QR code on this line can be obtained by calculating the distance of the QR code from the camera. This distance can be obtained by comparing the apparent width of the QR code in the image against the actual width, which is known.

Step 3: Once the location of a sufficient number of QR codes are obtained, we can solve for the configuration of the hand as follows.

![](https://github.com/janakagoon/kisa/blob/main/misc/images/index-finger-vectors-and-angles.png?raw=true)

Let's say we are trying to figure out the configuration of the index finger.

With reference to the figure above, this configuration is fully described by:

* The rotation of the wrist around to the y-axis in the world coordinate system: `THETA_wrist`,
* The rotation of the wrist around to the x-axis in the world coordinate system: 'PHI_wrist`, 
* How much the finger has spread out form the straight position:  `THETA_index_rot`
* The flex of the first joint of the index finger. This is the joint that connects it to the palm. `THETA_index_base`
* The flex of the middle joint, from the straight position relative to the previous segment. When the middle segment is parallel to the base segment this angle is zero.:  `THETA_index_middle`:
* The flex of joint closest to the tip: `THETA_index_tip`

To fully describe the configuration of the index finger, we need to solve for these six angles.

Now let us derive the positions of all the QR codes - whose locations are known in the World coordinate system - in terms of the angles and vectors described above.

For simplicity, let's assume the QR codes are located at the following locations:

* The QR code at the center of the wrist: `Q_wrist`
* The QR code at the point where the palm meets the finger, around which it rotates when spreading in and out: `Q_index_rot`
* The QR code at the base joint of the index finger: `Q_index_base`
* The QR code where the base segment meets the middle segment: `Q_index_middle`
* The QR code at where the middle segment meets the tip segment: `Q_index_tip`
* The QR code at the end of the tip segment: `Q_index_end`

Additionally the following 5 quantities are known from the dimensions of the hand.

* The vector from the center of the wrist `Q_wrist` to the center of rotation of the spread motion `Q_index_rot`:  `R^prime _wrist_to_rot`
* The vector from the center of rotation of the spread motion `Q_index_rot`, to the point to the base joint `Q_index_base`: `R^prime _rot_to_base`
* The length of the first base segment: `|R_base_to_middle|`
* The length of the middle finger segment: `|R_middle_to_tip|`
* The length of the tip finger segment: `|R_tip to end|`

Now we form the following equations which establishes a relationship between the above known quantities and the 6 unknown angles we need to solve for. 

Here R(PHI, THETA, RHO) is the rotation matrix that will rotate any vector around the x axis by PHI , around the y axis by THETA and around the z axis by RHO.

![](https://github.com/janakagoon/kisa/blob/main/misc/images/qr-code-solution-equations.png?raw=true)

As we have six equations with six scalar unknowns, the above equations will yield the configuration of the wrist and the index fingers.

These can be solved using one of many numerical linear algebra libraries that are available. 

## Appendix D: Investor Site

Interested in a pitch deck and the financials? Please  visit this site:

![](https://janakagoon.github.io/dexterousactuators/assets/image/DA-logo.png)

http://www.dexterousactuator.com


/


