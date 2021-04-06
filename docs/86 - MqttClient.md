# MQTT Client Libraries for Rust

>*Date Created: 04.04.2021*

**Documentation for [Story 86](https://softwaredefinedvehicle.visualstudio.com/SoftwareDefinedVehicle/_workitems/edit/86)**

The goal of this document is to examine the options available for MQTT client libraries in Rust. The following table shows key information from the most active and used options in the community.

| Name                          | GitHub URL                                                                             | MQTT        | Downloads | Contributors | Commits | Last Commit | License    | Notes                                                                                                                             |
| ----------------------------- | -------------------------------------------------------------------------------------- | ----------- | --------- | ------------ | ------- | ----------- | ---------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Eclipse Paho MQTT Rust Client | [https://github.com/eclipse/paho.mqtt.rust](https://github.com/eclipse/paho.mqtt.rust) | 3.1/3.1.1/5 | 24k       | 7            | 244     | 1/1/2021    | EPL        | This is an FFI wrapper around the Paho C MQTT library.  https://www.eclipse.org/paho/index.php?page=clients/rust/index.php        |
| Rumqtt                        | [https://github.com/bytebeamio/rumqtt](https://github.com/bytebeamio/rumqtt)           | 3.1.1/5     | 10k/20k   | 24           | 230     | 3/25/3021   | Apache-2.0 |                                                                                                                                   |
| Mqttrs                        | [https://github.com/00imvj00/mqttrs](https://github.com/00imvj00/mqttrs)               | 3.1.1       | 5k        | 6            | 132     | 1/10/2021   | Apache-2.0 |                                                                                                                                   |
| MQTT-rs                       | [https://github.com/zonyitoo/mqtt-rs](https://github.com/zonyitoo/mqtt-rs)             | 3.1.1       | 33k       | 11           | 130     | 3/28/2021   | MIT        |                                                                                                                                   |
| ntex-mqtt                     | [https://github.com/ntex-rs/ntex-mqtt](https://github.com/ntex-rs/ntex-mqtt)           | 3.1.1/5     | 4k        | 6\*          | 358     | 4/3/2021    | MIT        | Client/Server Framework. Based on ntex (https://docs.rs/ntex/0.3.6/ntex/).  \*Microsoft contributors: Nikolay Kim and Max Gortman |
| RustMQ                        | [https://github.com/inre/rust-mq](https://github.com/inre/rust-mq)                     | 3?          | 5k        | 10           | 127     | 12/9/2020   | MIT        |                                                                                                                                   |
|                               |                                                                                        |             |           |              |         |             |            |                                                                                                                                   |
| Mqtttest                      | [https://github.com/vincentdephily/mqttest](https://github.com/vincentdephily/mqttest) |             |           |              |         |             | Apache-2.0 | MQTT Server designed for unit testing clients.                                                                                    |

## MQTT Versions

The requirements around MQTT have not been defined and therefore it is unclear which version needs to be supported. The Azure IoT Hub currently [supports 3.1.1](https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-mqtt-support) and MQTT 5 support is in [preview](https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-mqtt-5). In addition, MQTT version support for the device broker will need to be considered once that is chosen.

The differences between 3.1 and 3.1.1 are documented [here](https://github.com/mqtt/mqtt.org/wiki/Differences-between-3.1.0-and-3.1.1) and those between 3.1.1 and 5.0 [here](https://github.com/mqtt/mqtt.org/wiki/Differences-between-3.1.1-and-5.0).

# Recommendations 

Given the lack of clear definition of requirements other than IoT Hub connectivity, my recommendation is that we look at the following 3 libraries: Eclipse Paho, Rumqtt and MQTT-rs.  The Eclipse Paho library is as close to an "official" Rust MQTT library as the community has, but given that it is simply a wrapper around the underlying Paho C library, the tradeoff against the safety and build simplicity of a pure Rust crate needs to be considered. Rumqtt and MQTT-rs are both native Rust libraries that have been around for a number of years and still see recent updates. One other factor is that these three libraries give us one of each license type, should concerns arise around a specific license.

