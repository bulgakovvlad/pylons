package types

func GenCoinInputList(name string, count int64) CoinInputList {
	return CoinInputList{
		CoinInput{
			Coin:  name,
			Count: count,
		},
	}
}

func GenItemInputList(alivePercent int, names ...string) ItemInputList {
	iiL := ItemInputList{}
	for _, name := range names {
		iiL = append(iiL, ItemInput{
			nil,
			nil,
			StringInputParamList{StringInputParam{"Name", name}},
			alivePercent,
			ItemUpgradeParams{},
		})
	}
	return iiL
}

func GenDetailedItemInputList(alivePercent int, itemUpgrades []ItemUpgradeParams, names ...string) ItemInputList {
	iiL := ItemInputList{}
	for idx, name := range names {
		toUpgrade := ItemUpgradeParams{}
		if idx < len(itemUpgrades) {
			toUpgrade = itemUpgrades[idx]
		}
		iiL = append(iiL, ItemInput{
			nil,
			nil,
			StringInputParamList{StringInputParam{"Name", name}},
			alivePercent,
			toUpgrade,
		})
	}
	return iiL
}

func GenCoinOnlyEntry(coinName string) EntriesList {
	return EntriesList{
		CoinOutput{
			Coin:  coinName,
			Count: 1,
		},
	}
}

func GenCoinOnlyEntryRand(coinName string) EntriesList {
	return EntriesList{
		CoinOutput{
			Coin:    coinName,
			Program: `randi(10)`,
		},
	}
}

func GenSingleItemInputList(alivePercent int, itemName string) ItemInputList {
	return ItemInputList{
		ItemInput{
			Doubles: DoubleInputParamList{},
			Longs:   LongInputParamList{},
			Strings: StringInputParamList{
				StringInputParam{
					Key:   "Name",
					Value: itemName,
				},
			},
			AlivePercent: alivePercent,
		},
	}
}

func GenItemNameUpgradeParams(desItemName string) ItemUpgradeParams {
	return ItemUpgradeParams{
		Doubles: DoubleParamList{},
		Longs:   LongParamList{},
		Strings: StringParamList{
			StringParam{
				Key:   "Name",
				Value: desItemName,
			},
		},
	}
}

func GenItemOnlyEntry(itemName string) EntriesList {
	return EntriesList{
		ItemOutput{
			Doubles: DoubleParamList{DoubleParam{Key: "endurance", DoubleWeightTable: DoubleWeightTable{WeightRanges: []DoubleWeightRange{
				DoubleWeightRange{
					Lower:  "100.00",
					Upper:  "500.00",
					Weight: 6,
				},
				DoubleWeightRange{
					Lower:  "501.00",
					Upper:  "800.00",
					Weight: 2,
				},
			}}, Rate: "1.0"}},
			Longs: LongParamList{LongParam{Key: "HP", IntWeightTable: IntWeightTable{WeightRanges: []IntWeightRange{
				IntWeightRange{
					Lower:  100,
					Upper:  500,
					Weight: 6,
				},
				IntWeightRange{
					Lower:  501,
					Upper:  800,
					Weight: 2,
				},
			}}}},
			Strings: StringParamList{StringParam{Key: "Name", Value: itemName, Rate: "1.0", Program: ""}},
		},
	}
}

func GenItemOnlyEntryRand(itemName string) EntriesList {
	return EntriesList{
		ItemOutput{
			Doubles: DoubleParamList{DoubleParam{
				Key:     "endurance",
				Program: `500.00`,
				Rate:    "1.0",
			}},
			Longs: LongParamList{LongParam{
				Key:     "HP",
				Program: `500 + randi(300)`,
				Rate:    "1.0",
			}},
			Strings: StringParamList{StringParam{Key: "Name", Value: itemName, Rate: "1.0", Program: ""}},
		},
	}
}

func GenOneOutput(n int) WeightedOutputsList {
	wol := WeightedOutputsList{}
	for i := 0; i < n; i++ {
		wol = append(wol, OutputsList{
			Result: []int{i},
			Weight: 1,
		})
	}
	return wol
}

func GenEntries(coinName string, itemName string) EntriesList {
	return EntriesList{
		GenCoinOnlyEntry(coinName)[0],
		GenItemOnlyEntry(itemName)[0],
	}
}

func GenEntriesRand(coinName, itemName string) EntriesList {
	return EntriesList{
		GenCoinOnlyEntryRand(coinName)[0],
		GenItemOnlyEntryRand(itemName)[0],
	}
}

func GenToUpgradeForString(targetKey, targetValue string) ItemUpgradeParams {
	return ItemUpgradeParams{
		Strings: StringParamList{
			{Key: targetKey, Value: targetValue},
		},
	}
}

func GenToUpgradeForLong(targetKey string, upgradeAmount int) ItemUpgradeParams {
	return ItemUpgradeParams{
		Longs: []LongParam{
			{
				Key: targetKey,
				IntWeightTable: IntWeightTable{WeightRanges: []IntWeightRange{
					IntWeightRange{
						Lower:  upgradeAmount,
						Upper:  upgradeAmount,
						Weight: 1,
					},
				}},
			},
		},
	}
}

func GenToUpgradeForDouble(targetKey string, upgradeAmount FloatString) ItemUpgradeParams {
	return ItemUpgradeParams{
		Doubles: []DoubleParam{
			{
				Key: targetKey,
				DoubleWeightTable: DoubleWeightTable{WeightRanges: []DoubleWeightRange{
					DoubleWeightRange{
						Lower:  upgradeAmount,
						Upper:  upgradeAmount,
						Weight: 1,
					},
				}},
			},
		},
	}
}
