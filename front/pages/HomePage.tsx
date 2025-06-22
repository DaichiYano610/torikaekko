import React from 'react';
import { View, Text, Button } from 'react-native';

export default function HomeScreen({}) {
  return (
    <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
      <Text>ホーム画面</Text>
      <Button
        title="詳細へ"
      />
    </View>
  );
}
