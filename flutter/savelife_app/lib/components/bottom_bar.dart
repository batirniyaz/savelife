import 'package:flutter/material.dart';

// Import your pages and styles here
import '../pages/home_page.dart';
import '../style/style.dart';

class BottomBar extends StatefulWidget {
  const BottomBar({Key? key}) : super(key: key);

  @override
  State<BottomBar> createState() => _BottomBarState();
}

class _BottomBarState extends State<BottomBar> {
  int _selectedIndex = 0;

  static const List<Widget> _widgetOptions = <Widget>[
    HomePage(),
    Text('Index 1: Business'),
    Text('Index 2: School'),
    Text('Index 3: School'),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(AppStyle.homeIcon),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(AppStyle.eventIcon),
            label: 'Schedule',
          ),
          BottomNavigationBarItem(
            icon: Icon(AppStyle.reportIcon),
            label: 'Report',
          ),
          BottomNavigationBarItem(
            icon: Icon(AppStyle.notificationIcon),
            label: 'Notifications',
          ),
        ],
        currentIndex: _selectedIndex,
        selectedItemColor: AppStyle.selectedColor,
        unselectedItemColor: AppStyle.unselectedColor,
        onTap: _onItemTapped,
      ),
    );
  }
}
