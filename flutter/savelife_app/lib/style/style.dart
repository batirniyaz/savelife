import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class AppStyle {
  // App Images ..
  static const String profile = "assets/images/profile.jpg";
  static const String image1 = "assets/images/image1.jpg";

  // App Icons.
  static const IconData findicon = Icons.find_in_page;
  static const IconData searchicon = Icons.search;

  // Define your colors
  static const Color selectedColor = Colors.red;
  static const Color unselectedColor = Colors.black; 

  // Bottom Bar Icons.
  static const IconData homeIcon = Icons.home;
  static const IconData eventIcon = Icons.event;
  static const IconData reportIcon = Icons.report;
  static const IconData notificationIcon = Icons.notifications;

  // App Colors
  static const primarySwatch = Colors.red;
  static const inputFillColor = Colors.blue;

  // App Theme Data..
  static ThemeData? theme = ThemeData(
      textTheme: GoogleFonts.nunitoSansTextTheme().apply(
    bodyColor: Colors.blueAccent,
    displayColor: Colors.cyan,
  ),);
}
