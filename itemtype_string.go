// Code generated by "stringer -type=itemType"; DO NOT EDIT.

package yql_elastic

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[itemTerm-0]
	_ = x[itemLeftGroupDelim-1]
	_ = x[itemRightGroupDelim-2]
	_ = x[itemField-3]
	_ = x[itemFieldValue-4]
	_ = x[itemMust-5]
	_ = x[itemMustNot-6]
	_ = x[itemSkipWhitespace-7]
	_ = x[itemLowerThan-8]
	_ = x[itemGreaterThan-9]
	_ = x[itemKeyword-10]
	_ = x[itemRegex-11]
}

const _itemType_name = "itemTermitemLeftGroupDelimitemRightGroupDelimitemFielditemFieldValueitemMustitemMustNotitemSkipWhitespaceitemLowerThanitemGreaterThanitemKeyworditemRegex"

var _itemType_index = [...]uint8{0, 8, 26, 45, 54, 68, 76, 87, 105, 118, 133, 144, 153}

func (i itemType) String() string {
	if i < 0 || i >= itemType(len(_itemType_index)-1) {
		return "itemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}