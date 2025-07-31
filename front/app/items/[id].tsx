import React from 'react';
import { View, Text } from 'react-native';
import { useLocalSearchParams } from 'expo-router';

export default function ItemDetailScreen() {
  const { id } = useLocalSearchParams();

  // 通常はここでidを使ってAPIから詳細データを取得する
  return (
    <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
      <Text style={{ fontSize: 24 }}>二谷耕太郎の詳細</Text>
      <Text style={{ fontSize: 18, marginTop: 12 }}>ID: {id}</Text>
    </View>
  );
}
