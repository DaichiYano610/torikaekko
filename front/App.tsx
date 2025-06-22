import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import HomePage from "./pages/HomePage"
const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <NavigationContainer>
j     <Stack.Navigator initialRouteName="Home">
        <Stack.Screen name="Home" component={HomePage} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}

