import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

class Services {
  final String image;
  final Color color;
  Services({
    required this.image,
    required this.color,
  });
}

// Services List of Data

List<Services> servicesList = [
  Services(
    image: "assets/icons/icons8-здоровье-64.png",
    color: Colors.pink,
  ),
  Services(
    image: "assets/icons/icons8-medical-test-66.png",
    color: Colors.amber,
  ),
  Services(
    image: "assets/images/noun-medical-test.jpg",
    color: const Color(0xffD6F6FF),
  ),
  Services(
    image: "assets/images/noun-corona.jpg",
    color: const Color(0xffF2E3E9),
  ),
];

/// Upcoming Appointments

class Appointments {
  final String date;
  final String time;
  final String title;
  final String subTitle;
  final Color color;

  Appointments({
    required this.color,
    required this.date,
    required this.time,
    required this.title,
    required this.subTitle,
  });
}

// Upcoming Appoinments List of Data

List<Appointments> upcomingAppointmentsList = [
  Appointments(
    date: "12\nTue",
    time: DateFormat('hh:mm a').format(DateTime.now()),
    title: "Dr. Shaynyyy",
    subTitle: "Depression",
    color: const Color(0xff1C6BA4),
  ),
  Appointments(
    date: "12\nTue",
    time: DateFormat('hh:mm a').format(DateTime.now()),
    title: "Dr. Batya",
    subTitle: "Depression",
    color: const Color(0xff1C6BA4),
  ),
];
