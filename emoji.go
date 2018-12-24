package insights

import (
	"strings"
)

var emojis = []string{
	"😀",
	"😁",
	"😂",
	"🤣",
	"😃",
	"😄",
	"😅",
	"😆",
	"😉",
	"😊",
	"😋",
	"😎",
	"😍",
	"😘",
	"🥰",
	"😗",
	"😙",
	"😚",
	"☺️",
	"🙂",
	"🤗",
	"🤩",
	"🤔",
	"🤨",
	"😐",
	"😑",
	"😶",
	"🙄",
	"😏",
	"😣",
	"😥",
	"😮",
	"🤐",
	"😯",
	"😪",
	"😫",
	"😴",
	"😌",
	"😛",
	"😜",
	"😝",
	"🤤",
	"😒",
	"😓",
	"😔",
	"😕",
	"🙃",
	"🤑",
	"😲",
	"☹️",
	"🙁",
	"😖",
	"😞",
	"😟",
	"😤",
	"😢",
	"😭",
	"😦",
	"😧",
	"😨",
	"😩",
	"🤯",
	"😬",
	"😰",
	"😱",
	"🥵",
	"🥶",
	"😳",
	"🤪",
	"😵",
	"😡",
	"😠",
	"🤬",
	"😷",
	"🤒",
	"🤕",
	"🤢",
	"🤮",
	"🤧",
	"😇",
	"🤠",
	"🤡",
	"🥳",
	"🥴",
	"🥺",
	"🤥",
	"🤫",
	"🤭",
	"🧐",
	"🤓",
	"😈",
	"👿",
	"👹",
	"👺",
	"💀",
	"👻",
	"👽",
	"🤖",
	"💩",
	"😺",
	"😸",
	"😹",
	"😻",
	"😼",
	"😽",
	"🙀",
	"😿",
	"😾",
	"👶",
	"👧",
	"🧒",
	"👦",
	"👩",
	"🧑",
	"👨",
	"👵",
	"🧓",
	"👴",
	"👲",
	"👳‍♀️",
	"👳‍♂️",
	"🧕",
	"🧔",
	"👱‍♂️",
	"👱‍♀️",
	"👨‍🦰",
	"👩‍🦰",
	"👨‍🦱",
	"👩‍🦱",
	"👨‍🦲",
	"👩‍🦲",
	"👨‍🦳",
	"👩‍🦳",
	"🦸‍♀️",
	"🦸‍♂️",
	"🦹‍♀️",
	"🦹‍♂️",
	"👮‍♀️",
	"👮‍♂️",
	"👷‍♀️",
	"👷‍♂️",
	"💂‍♀️",
	"💂‍♂️",
	"🕵️‍♀️",
	"🕵️‍♂️",
	"👩‍⚕️",
	"👨‍⚕️",
	"👩‍🌾",
	"👨‍🌾",
	"👩‍🍳",
	"👨‍🍳",
	"👩‍🎓",
	"👨‍🎓",
	"👩‍🎤",
	"👨‍🎤",
	"👩‍🏫",
	"👨‍🏫",
	"👩‍🏭",
	"👨‍🏭",
	"👩‍💻",
	"👨‍💻",
	"👩‍💼",
	"👨‍💼",
	"👩‍🔧",
	"👨‍🔧",
	"👩‍🔬",
	"👨‍🔬",
	"👩‍🎨",
	"👨‍🎨",
	"👩‍🚒",
	"👨‍🚒",
	"👩‍✈️",
	"👨‍✈️",
	"👩‍🚀",
	"👨‍🚀",
	"👩‍⚖️",
	"👨‍⚖️",
	"👰",
	"🤵",
	"👸",
	"🤴",
	"🤶",
	"🎅",
	"🧙‍♀️",
	"🧙‍♂️",
	"🧝‍♀️",
	"🧝‍♂️",
	"🧛‍♀️",
	"🧛‍♂️",
	"🧟‍♀️",
	"🧟‍♂️",
	"🧞‍♀️",
	"🧞‍♂️",
	"🧜‍♀️",
	"🧜‍♂️",
	"🧚‍♀️",
	"🧚‍♂️",
	"👼",
	"🤰",
	"🤱",
	"🙇‍♀️",
	"🙇‍♂️",
	"💁‍♀️",
	"💁‍♂️",
	"🙅‍♀️",
	"🙅‍♂️",
	"🙆‍♀️",
	"🙆‍♂️",
	"🙋‍♀️",
	"🙋‍♂️",
	"🤦‍♀️",
	"🤦‍♂️",
	"🤷‍♀️",
	"🤷‍♂️",
	"🙎‍♀️",
	"🙎‍♂️",
	"🙍‍♀️",
	"🙍‍♂️",
	"💇‍♀️",
	"💇‍♂️",
	"💆‍♀️",
	"💆‍♂️",
	"🧖‍♀️",
	"🧖‍♂️",
	"💅",
	"🤳",
	"💃",
	"🕺",
	"👯‍♀️",
	"👯‍♂️",
	"🕴",
	"🚶‍♀️",
	"🚶‍♂️",
	"🏃‍♀️",
	"🏃‍♂️",
	"👫",
	"👭",
	"👬",
	"💑",
	"👩‍❤️‍👩",
	"👨‍❤️‍👨",
	"💏",
	"👩‍❤️‍💋‍👩",
	"👨‍❤️‍💋‍👨",
	"👪",
	"👨‍👩‍👧",
	"👨‍👩‍👧‍👦",
	"👨‍👩‍👦‍👦",
	"👨‍👩‍👧‍👧",
	"👩‍👩‍👦",
	"👩‍👩‍👧",
	"👩‍👩‍👧‍👦",
	"👩‍👩‍👦‍👦",
	"👩‍👩‍👧‍👧",
	"👨‍👨‍👦",
	"👨‍👨‍👧",
	"👨‍👨‍👧‍👦",
	"👨‍👨‍👦‍👦",
	"👨‍👨‍👧‍👧",
	"👩‍👦",
	"👩‍👧",
	"👩‍👧‍👦",
	"👩‍👦‍👦",
	"👩‍👧‍👧",
	"👨‍👦",
	"👨‍👧",
	"👨‍👧‍👦",
	"👨‍👦‍👦",
	"👨‍👧‍👧",
	"🤲",
	"👐",
	"🙌",
	"👏",
	"🤝",
	"👍",
	"👎",
	"👊",
	"✊",
	"🤛",
	"🤜",
	"🤞",
	"✌️",
	"🤟",
	"🤘",
	"👌",
	"👈",
	"👉",
	"👆",
	"👇",
	"☝️",
	"✋",
	"🤚",
	"🖐",
	"🖖",
	"👋",
	"🤙",
	"💪",
	"🦵",
	"🦶",
	"🖕",
	"✍️",
	"🙏",
	"💍",
	"💄",
	"💋",
	"👄",
	"👅",
	"👂",
	"👃",
	"👣",
	"👁",
	"👀",
	"🧠",
	"🦴",
	"🦷",
	"🗣",
	"👤",
	"👥",
	"🧥",
	"👚",
	"👕",
	"👖",
	"👔",
	"👗",
	"👙",
	"👘",
	"👠",
	"👡",
	"👢",
	"👞",
	"👟",
	"🥾",
	"🥿",
	"🧦",
	"🧤",
	"🧣",
	"🎩",
	"🧢",
	"👒",
	"🎓",
	"⛑",
	"👑",
	"👝",
	"👛",
	"👜",
	"💼",
	"🎒",
	"👓",
	"🕶",
	"🥽",
	"🥼",
	"🌂",
	"🧵",
	"🧶",
	"👶🏻",
	"👦🏻",
	"👧🏻",
	"👨🏻",
	"👩🏻",
	"👱🏻‍♀️",
	"👱🏻",
	"👴🏻",
	"👵🏻",
	"👲🏻",
	"👳🏻‍♀️",
	"👳🏻",
	"👮🏻‍♀️",
	"👮🏻",
	"👷🏻‍♀️",
	"👷🏻",
	"💂🏻‍♀️",
	"💂🏻",
	"🕵🏻‍♀️",
	"🕵🏻",
	"👩🏻‍⚕️",
	"👨🏻‍⚕️",
	"👩🏻‍🌾",
	"👨🏻‍🌾",
	"👩🏻‍🍳",
	"👨🏻‍🍳",
	"👩🏻‍🎓",
	"👨🏻‍🎓",
	"👩🏻‍🎤",
	"👨🏻‍🎤",
	"👩🏻‍🏫",
	"👨🏻‍🏫",
	"👩🏻‍🏭",
	"👨🏻‍🏭",
	"👩🏻‍💻",
	"👨🏻‍💻",
	"👩🏻‍💼",
	"👨🏻‍💼",
	"👩🏻‍🔧",
	"👨🏻‍🔧",
	"👩🏻‍🔬",
	"👨🏻‍🔬",
	"👩🏻‍🎨",
	"👨🏻‍🎨",
	"👩🏻‍🚒",
	"👨🏻‍🚒",
	"👩🏻‍✈️",
	"👨🏻‍✈️",
	"👩🏻‍🚀",
	"👨🏻‍🚀",
	"👩🏻‍⚖️",
	"👨🏻‍⚖️",
	"🤶🏻",
	"🎅🏻",
	"👸🏻",
	"🤴🏻",
	"👰🏻",
	"🤵🏻",
	"👼🏻",
	"🤰🏻",
	"🙇🏻‍♀️",
	"🙇🏻",
	"💁🏻",
	"💁🏻‍♂️",
	"🙅🏻",
	"🙅🏻‍♂️",
	"🙆🏻",
	"🙆🏻‍♂️",
	"🙋🏻",
	"🙋🏻‍♂️",
	"🤦🏻‍♀️",
	"🤦🏻‍♂️",
	"🤷🏻‍♀️",
	"🤷🏻‍♂️",
	"🙎🏻",
	"🙎🏻‍♂️",
	"🙍🏻",
	"🙍🏻‍♂️",
	"💇🏻",
	"💇🏻‍♂️",
	"💆🏻",
	"💆🏻‍♂️",
	"🕴🏻",
	"💃🏻",
	"🕺🏻",
	"🚶🏻‍♀️",
	"🚶🏻",
	"🏃🏻‍♀️",
	"🏃🏻",
	"🤲🏻",
	"👐🏻",
	"🙌🏻",
	"👏🏻",
	"🙏🏻",
	"👍🏻",
	"👎🏻",
	"👊🏻",
	"✊🏻",
	"🤛🏻",
	"🤜🏻",
	"🤞🏻",
	"✌🏻",
	"🤟🏻",
	"🤘🏻",
	"👌🏻",
	"👈🏻",
	"👉🏻",
	"👆🏻",
	"👇🏻",
	"☝🏻",
	"✋🏻",
	"🤚🏻",
	"🖐🏻",
	"🖖🏻",
	"👋🏻",
	"🤙🏻",
	"💪🏻",
	"🖕🏻",
	"✍🏻",
	"🤳🏻",
	"💅🏻",
	"👂🏻",
	"👃🏻",
	"👶🏼",
	"👦🏼",
	"👧🏼",
	"👨🏼",
	"👩🏼",
	"👱🏼‍♀️",
	"👱🏼",
	"👴🏼",
	"👵🏼",
	"👲🏼",
	"👳🏼‍♀️",
	"👳🏼",
	"👮🏼‍♀️",
	"👮🏼",
	"👷🏼‍♀️",
	"👷🏼",
	"💂🏼‍♀️",
	"💂🏼",
	"🕵🏼‍♀️",
	"🕵🏼",
	"👩🏼‍⚕️",
	"👨🏼‍⚕️",
	"👩🏼‍🌾",
	"👨🏼‍🌾",
	"👩🏼‍🍳",
	"👨🏼‍🍳",
	"👩🏼‍🎓",
	"👨🏼‍🎓",
	"👩🏼‍🎤",
	"👨🏼‍🎤",
	"👩🏼‍🏫",
	"👨🏼‍🏫",
	"👩🏼‍🏭",
	"👨🏼‍🏭",
	"👩🏼‍💻",
	"👨🏼‍💻",
	"👩🏼‍💼",
	"👨🏼‍💼",
	"👩🏼‍🔧",
	"👨🏼‍🔧",
	"👩🏼‍🔬",
	"👨🏼‍🔬",
	"👩🏼‍🎨",
	"👨🏼‍🎨",
	"👩🏼‍🚒",
	"👨🏼‍🚒",
	"👩🏼‍✈️",
	"👨🏼‍✈️",
	"👩🏼‍🚀",
	"👨🏼‍🚀",
	"👩🏼‍⚖️",
	"👨🏼‍⚖️",
	"🤶🏼",
	"🎅🏼",
	"👸🏼",
	"🤴🏼",
	"👰🏼",
	"🤵🏼",
	"👼🏼",
	"🤰🏼",
	"🙇🏼‍♀️",
	"🙇🏼",
	"💁🏼",
	"💁🏼‍♂️",
	"🙅🏼",
	"🙅🏼‍♂️",
	"🙆🏼",
	"🙆🏼‍♂️",
	"🙋🏼",
	"🙋🏼‍♂️",
	"🤦🏼‍♀️",
	"🤦🏼‍♂️",
	"🤷🏼‍♀️",
	"🤷🏼‍♂️",
	"🙎🏼",
	"🙎🏼‍♂️",
	"🙍🏼",
	"🙍🏼‍♂️",
	"💇🏼",
	"💇🏼‍♂️",
	"💆🏼",
	"💆🏼‍♂️",
	"🕴🏼",
	"💃🏼",
	"🕺🏼",
	"🚶🏼‍♀️",
	"🚶🏼",
	"🏃🏼‍♀️",
	"🏃🏼",
	"🤲🏼",
	"👐🏼",
	"🙌🏼",
	"👏🏼",
	"🙏🏼",
	"👍🏼",
	"👎🏼",
	"👊🏼",
	"✊🏼",
	"🤛🏼",
	"🤜🏼",
	"🤞🏼",
	"✌🏼",
	"🤟🏼",
	"🤘🏼",
	"👌🏼",
	"👈🏼",
	"👉🏼",
	"👆🏼",
	"👇🏼",
	"☝🏼",
	"✋🏼",
	"🤚🏼",
	"🖐🏼",
	"🖖🏼",
	"👋🏼",
	"🤙🏼",
	"💪🏼",
	"🖕🏼",
	"✍🏼",
	"🤳🏼",
	"💅🏼",
	"👂🏼",
	"👃🏼",
	"👶🏽",
	"👦🏽",
	"👧🏽",
	"👨🏽",
	"👩🏽",
	"👱🏽‍♀️",
	"👱🏽",
	"👴🏽",
	"👵🏽",
	"👲🏽",
	"👳🏽‍♀️",
	"👳🏽",
	"👮🏽‍♀️",
	"👮🏽",
	"👷🏽‍♀️",
	"👷🏽",
	"💂🏽‍♀️",
	"💂🏽",
	"🕵🏽‍♀️",
	"🕵🏽",
	"👩🏽‍⚕️",
	"👨🏽‍⚕️",
	"👩🏽‍🌾",
	"👨🏽‍🌾",
	"👩🏽‍🍳",
	"👨🏽‍🍳",
	"👩🏽‍🎓",
	"👨🏽‍🎓",
	"👩🏽‍🎤",
	"👨🏽‍🎤",
	"👩🏽‍🏫",
	"👨🏽‍🏫",
	"👩🏽‍🏭",
	"👨🏽‍🏭",
	"👩🏽‍💻",
	"👨🏽‍💻",
	"👩🏽‍💼",
	"👨🏽‍💼",
	"👩🏽‍🔧",
	"👨🏽‍🔧",
	"👩🏽‍🔬",
	"👨🏽‍🔬",
	"👩🏽‍🎨",
	"👨🏽‍🎨",
	"👩🏽‍🚒",
	"👨🏽‍🚒",
	"👩🏽‍✈️",
	"👨🏽‍✈️",
	"👩🏽‍🚀",
	"👨🏽‍🚀",
	"👩🏽‍⚖️",
	"👨🏽‍⚖️",
	"🤶🏽",
	"🎅🏽",
	"👸🏽",
	"🤴🏽",
	"👰🏽",
	"🤵🏽",
	"👼🏽",
	"🤰🏽",
	"🙇🏽‍♀️",
	"🙇🏽",
	"💁🏽",
	"💁🏽‍♂️",
	"🙅🏽",
	"🙅🏽‍♂️",
	"🙆🏽",
	"🙆🏽‍♂️",
	"🙋🏽",
	"🙋🏽‍♂️",
	"🤦🏽‍♀️",
	"🤦🏽‍♂️",
	"🤷🏽‍♀️",
	"🤷🏽‍♂️",
	"🙎🏽",
	"🙎🏽‍♂️",
	"🙍🏽",
	"🙍🏽‍♂️",
	"💇🏽",
	"💇🏽‍♂️",
	"💆🏽",
	"💆🏽‍♂️",
	"🕴🏼",
	"💃🏽",
	"🕺🏽",
	"🚶🏽‍♀️",
	"🚶🏽",
	"🏃🏽‍♀️",
	"🏃🏽",
	"🤲🏽",
	"👐🏽",
	"🙌🏽",
	"👏🏽",
	"🙏🏽",
	"👍🏽",
	"👎🏽",
	"👊🏽",
	"✊🏽",
	"🤛🏽",
	"🤜🏽",
	"🤞🏽",
	"✌🏽",
	"🤟🏽",
	"🤘🏽",
	"👌🏽",
	"👈🏽",
	"👉🏽",
	"👆🏽",
	"👇🏽",
	"☝🏽",
	"✋🏽",
	"🤚🏽",
	"🖐🏽",
	"🖖🏽",
	"👋🏽",
	"🤙🏽",
	"💪🏽",
	"🖕🏽",
	"✍🏽",
	"🤳🏽",
	"💅🏽",
	"👂🏽",
	"👃🏽",
	"👶🏾",
	"👦🏾",
	"👧🏾",
	"👨🏾",
	"👩🏾",
	"👱🏾‍♀️",
	"👱🏾",
	"👴🏾",
	"👵🏾",
	"👲🏾",
	"👳🏾‍♀️",
	"👳🏾",
	"👮🏾‍♀️",
	"👮🏾",
	"👷🏾‍♀️",
	"👷🏾",
	"💂🏾‍♀️",
	"💂🏾",
	"🕵🏾‍♀️",
	"🕵🏾",
	"👩🏾‍⚕️",
	"👨🏾‍⚕️",
	"👩🏾‍🌾",
	"👨🏾‍🌾",
	"👩🏾‍🍳",
	"👨🏾‍🍳",
	"👩🏾‍🎓",
	"👨🏾‍🎓",
	"👩🏾‍🎤",
	"👨🏾‍🎤",
	"👩🏾‍🏫",
	"👨🏾‍🏫",
	"👩🏾‍🏭",
	"👨🏾‍🏭",
	"👩🏾‍💻",
	"👨🏾‍💻",
	"👩🏾‍💼",
	"👨🏾‍💼",
	"👩🏾‍🔧",
	"👨🏾‍🔧",
	"👩🏾‍🔬",
	"👨🏾‍🔬",
	"👩🏾‍🎨",
	"👨🏾‍🎨",
	"👩🏾‍🚒",
	"👨🏾‍🚒",
	"👩🏾‍✈️",
	"👨🏾‍✈️",
	"👩🏾‍🚀",
	"👨🏾‍🚀",
	"👩🏾‍⚖️",
	"👨🏾‍⚖️",
	"🤶🏾",
	"🎅🏾",
	"👸🏾",
	"🤴🏾",
	"👰🏾",
	"🤵🏾",
	"👼🏾",
	"🤰🏾",
	"🙇🏾‍♀️",
	"🙇🏾",
	"💁🏾",
	"💁🏾‍♂️",
	"🙅🏾",
	"🙅🏾‍♂️",
	"🙆🏾",
	"🙆🏾‍♂️",
	"🙋🏾",
	"🙋🏾‍♂️",
	"🤦🏾‍♀️",
	"🤦🏾‍♂️",
	"🤷🏾‍♀️",
	"🤷🏾‍♂️",
	"🙎🏾",
	"🙎🏾‍♂️",
	"🙍🏾",
	"🙍🏾‍♂️",
	"💇🏾",
	"💇🏾‍♂️",
	"💆🏾",
	"💆🏾‍♂️",
	"🕴🏾",
	"💃🏾",
	"🕺🏾",
	"🚶🏾‍♀️",
	"🚶🏾",
	"🏃🏾‍♀️",
	"🏃🏾",
	"🤲🏾",
	"👐🏾",
	"🙌🏾",
	"👏🏾",
	"🙏🏾",
	"👍🏾",
	"👎🏾",
	"👊🏾",
	"✊🏾",
	"🤛🏾",
	"🤜🏾",
	"🤞🏾",
	"✌🏾",
	"🤟🏾",
	"🤘🏾",
	"👌🏾",
	"👈🏾",
	"👉🏾",
	"👆🏾",
	"👇🏾",
	"☝🏾",
	"✋🏾",
	"🤚🏾",
	"🖐🏾",
	"🖖🏾",
	"👋🏾",
	"🤙🏾",
	"💪🏾",
	"🖕🏾",
	"✍🏾",
	"🤳🏾",
	"💅🏾",
	"👂🏾",
	"👃🏾",
	"👶🏿",
	"👦🏿",
	"👧🏿",
	"👨🏿",
	"👩🏿",
	"👱🏿‍♀️",
	"👱🏿",
	"👴🏿",
	"👵🏿",
	"👲🏿",
	"👳🏿‍♀️",
	"👳🏿",
	"👮🏿‍♀️",
	"👮🏿",
	"👷🏿‍♀️",
	"👷🏿",
	"💂🏿‍♀️",
	"💂🏿",
	"🕵🏿‍♀️",
	"🕵🏿",
	"👩🏿‍⚕️",
	"👨🏿‍⚕️",
	"👩🏿‍🌾",
	"👨🏿‍🌾",
	"👩🏿‍🍳",
	"👨🏿‍🍳",
	"👩🏿‍🎓",
	"👨🏿‍🎓",
	"👩🏿‍🎤",
	"👨🏿‍🎤",
	"👩🏿‍🏫",
	"👨🏿‍🏫",
	"👩🏿‍🏭",
	"👨🏿‍🏭",
	"👩🏿‍💻",
	"👨🏿‍💻",
	"👩🏿‍💼",
	"👨🏿‍💼",
	"👩🏿‍🔧",
	"👨🏿‍🔧",
	"👩🏿‍🔬",
	"👨🏿‍🔬",
	"👩🏿‍🎨",
	"👨🏿‍🎨",
	"👩🏿‍🚒",
	"👨🏿‍🚒",
	"👩🏿‍✈️",
	"👨🏿‍✈️",
	"👩🏿‍🚀",
	"👨🏿‍🚀",
	"👩🏿‍⚖️",
	"👨🏿‍⚖️",
	"🤶🏿",
	"🎅🏿",
	"👸🏿",
	"🤴🏿",
	"👰🏿",
	"🤵🏿",
	"👼🏿",
	"🤰🏿",
	"🙇🏿‍♀️",
	"🙇🏿",
	"💁🏿",
	"💁🏿‍♂️",
	"🙅🏿",
	"🙅🏿‍♂️",
	"🙆🏿",
	"🙆🏿‍♂️",
	"🙋🏿",
	"🙋🏿‍♂️",
	"🤦🏿‍♀️",
	"🤦🏿‍♂️",
	"🤷🏿‍♀️",
	"🤷🏿‍♂️",
	"🙎🏿",
	"🙎🏿‍♂️",
	"🙍🏿",
	"🙍🏿‍♂️",
	"💇🏿",
	"💇🏿‍♂️",
	"💆🏿",
	"💆🏿‍♂️",
	"🕴🏿",
	"💃🏿",
	"🕺🏿",
	"🚶🏿‍♀️",
	"🚶🏿",
	"🏃🏿‍♀️",
	"🏃🏿",
	"🤲🏿",
	"👐🏿",
	"🙌🏿",
	"👏🏿",
	"🙏🏿",
	"👍🏿",
	"👎🏿",
	"👊🏿",
	"✊🏿",
	"🤛🏿",
	"🤜🏿",
	"🤞🏿",
	"✌🏿",
	"🤟🏿",
	"🤘🏿",
	"👌🏿",
	"👈🏿",
	"👉🏿",
	"👆🏿",
	"👇🏿",
	"☝🏿",
	"✋🏿",
	"🤚🏿",
	"🖐🏿",
	"🖖🏿",
	"👋🏿",
	"🤙🏿",
	"💪🏿",
	"🖕🏿",
	"✍🏿",
	"🤳🏿",
	"💅🏿",
	"👂🏿",
	"👃🏿🐶",
	"🐱",
	"🐭",
	"🐹",
	"🐰",
	"🦊",
	"🦝",
	"🐻",
	"🐼",
	"🦘",
	"🦡",
	"🐨",
	"🐯",
	"🦁",
	"🐮",
	"🐷",
	"🐽",
	"🐸",
	"🐵",
	"🙈",
	"🙉",
	"🙊",
	"🐒",
	"🐔",
	"🐧",
	"🐦",
	"🐤",
	"🐣",
	"🐥",
	"🦆",
	"🦢",
	"🦅",
	"🦉",
	"🦚",
	"🦜",
	"🦇",
	"🐺",
	"🐗",
	"🐴",
	"🦄",
	"🐝",
	"🐛",
	"🦋",
	"🐌",
	"🐚",
	"🐞",
	"🐜",
	"🦗",
	"🕷",
	"🕸",
	"🦂",
	"🦟",
	"🦠",
	"🐢",
	"🐍",
	"🦎",
	"🦖",
	"🦕",
	"🐙",
	"🦑",
	"🦐",
	"🦀",
	"🐡",
	"🐠",
	"🐟",
	"🐬",
	"🐳",
	"🐋",
	"🦈",
	"🐊",
	"🐅",
	"🐆",
	"🦓",
	"🦍",
	"🐘",
	"🦏",
	"🦛",
	"🐪",
	"🐫",
	"🦙",
	"🦒",
	"🐃",
	"🐂",
	"🐄",
	"🐎",
	"🐖",
	"🐏",
	"🐑",
	"🐐",
	"🦌",
	"🐕",
	"🐩",
	"🐈",
	"🐓",
	"🦃",
	"🕊",
	"🐇",
	"🐁",
	"🐀",
	"🐿",
	"🦔",
	"🐾",
	"🐉",
	"🐲",
	"🌵",
	"🎄",
	"🌲",
	"🌳",
	"🌴",
	"🌱",
	"🌿",
	"☘️",
	"🍀",
	"🎍",
	"🎋",
	"🍃",
	"🍂",
	"🍁",
	"🍄",
	"🌾",
	"💐",
	"🌷",
	"🌹",
	"🥀",
	"🌺",
	"🌸",
	"🌼",
	"🌻",
	"🌞",
	"🌝",
	"🌛",
	"🌜",
	"🌚",
	"🌕",
	"🌖",
	"🌗",
	"🌘",
	"🌑",
	"🌒",
	"🌓",
	"🌔",
	"🌙",
	"🌎",
	"🌍",
	"🌏",
	"💫",
	"⭐️",
	"🌟",
	"✨",
	"⚡️",
	"☄️",
	"💥",
	"🔥",
	"🌪",
	"🌈",
	"☀️",
	"🌤",
	"⛅️",
	"🌥",
	"☁️",
	"🌦",
	"🌧",
	"⛈",
	"🌩",
	"🌨",
	"❄️",
	"☃️",
	"⛄️",
	"🌬",
	"💨",
	"💧",
	"💦",
	"☔️",
	"☂️",
	"🌊",
	"🌫",
	"🍏",
	"🍎",
	"🍐",
	"🍊",
	"🍋",
	"🍌",
	"🍉",
	"🍇",
	"🍓",
	"🍈",
	"🍒",
	"🍑",
	"🍍",
	"🥭",
	"🥥",
	"🥝",
	"🍅",
	"🍆",
	"🥑",
	"🥦",
	"🥒",
	"🥬",
	"🌶",
	"🌽",
	"🥕",
	"🥔",
	"🍠",
	"🥐",
	"🍞",
	"🥖",
	"🥨",
	"🥯",
	"🧀",
	"🥚",
	"🍳",
	"🥞",
	"🥓",
	"🥩",
	"🍗",
	"🍖",
	"🌭",
	"🍔",
	"🍟",
	"🍕",
	"🥪",
	"🥙",
	"🌮",
	"🌯",
	"🥗",
	"🥘",
	"🥫",
	"🍝",
	"🍜",
	"🍲",
	"🍛",
	"🍣",
	"🍱",
	"🥟",
	"🍤",
	"🍙",
	"🍚",
	"🍘",
	"🍥",
	"🥮",
	"🥠",
	"🍢",
	"🍡",
	"🍧",
	"🍨",
	"🍦",
	"🥧",
	"🍰",
	"🎂",
	"🍮",
	"🍭",
	"🍬",
	"🍫",
	"🍿",
	"🧂",
	"🍩",
	"🍪",
	"🌰",
	"🥜",
	"🍯",
	"🥛",
	"🍼",
	"☕️",
	"🍵",
	"🥤",
	"🍶",
	"🍺",
	"🍻",
	"🥂",
	"🍷",
	"🥃",
	"🍸",
	"🍹",
	"🍾",
	"🥄",
	"🍴",
	"🍽",
	"🥣",
	"🥡",
	"🥢",
	"⚽️",
	"🏀",
	"🏈",
	"⚾️",
	"🥎",
	"🏐",
	"🏉",
	"🎾",
	"🥏",
	"🎱",
	"🏓",
	"🏸",
	"🥅",
	"🏒",
	"🏑",
	"🥍",
	"🏏",
	"⛳️",
	"🏹",
	"🎣",
	"🥊",
	"🥋",
	"🎽",
	"⛸",
	"🥌",
	"🛷",
	"🛹",
	"🎿",
	"⛷",
	"🏂",
	"🏋️‍♀️",
	"🏋🏻‍♀️",
	"🏋🏼‍♀️",
	"🏋🏽‍♀️",
	"🏋🏾‍♀️",
	"🏋🏿‍♀️",
	"🏋️‍♂️",
	"🏋🏻‍♂️",
	"🏋🏼‍♂️",
	"🏋🏽‍♂️",
	"🏋🏾‍♂️",
	"🏋🏿‍♂️",
	"🤼‍♀️",
	"🤼‍♂️",
	"🤸‍♀️",
	"🤸🏻‍♀️",
	"🤸🏼‍♀️",
	"🤸🏽‍♀️",
	"🤸🏾‍♀️",
	"🤸🏿‍♀️",
	"🤸‍♂️",
	"🤸🏻‍♂️",
	"🤸🏼‍♂️",
	"🤸🏽‍♂️",
	"🤸🏾‍♂️",
	"🤸🏿‍♂️",
	"⛹️‍♀️",
	"⛹🏻‍♀️",
	"⛹🏼‍♀️",
	"⛹🏽‍♀️",
	"⛹🏾‍♀️",
	"⛹🏿‍♀️",
	"⛹️‍♂️",
	"⛹🏻‍♂️",
	"⛹🏼‍♂️",
	"⛹🏽‍♂️",
	"⛹🏾‍♂️",
	"⛹🏿‍♂️",
	"🤺",
	"🤾‍♀️",
	"🤾🏻‍♀️",
	"🤾🏼‍♀️",
	"🤾🏾‍♀️",
	"🤾🏾‍♀️",
	"🤾🏿‍♀️",
	"🤾‍♂️",
	"🤾🏻‍♂️",
	"🤾🏼‍♂️",
	"🤾🏽‍♂️",
	"🤾🏾‍♂️",
	"🤾🏿‍♂️",
	"🏌️‍♀️",
	"🏌🏻‍♀️",
	"🏌🏼‍♀️",
	"🏌🏽‍♀️",
	"🏌🏾‍♀️",
	"🏌🏿‍♀️",
	"🏌️‍♂️",
	"🏌🏻‍♂️",
	"🏌🏼‍♂️",
	"🏌🏽‍♂️",
	"🏌🏾‍♂️",
	"🏌🏿‍♂️",
	"🏇",
	"🏇🏻",
	"🏇🏼",
	"🏇🏽",
	"🏇🏾",
	"🏇🏿",
	"🧘‍♀️",
	"🧘🏻‍♀️",
	"🧘🏼‍♀️",
	"🧘🏽‍♀️",
	"🧘🏾‍♀️",
	"🧘🏿‍♀️",
	"🧘‍♂️",
	"🧘🏻‍♂️",
	"🧘🏼‍♂️",
	"🧘🏽‍♂️",
	"🧘🏾‍♂️",
	"🧘🏿‍♂️",
	"🏄‍♀️",
	"🏄🏻‍♀️",
	"🏄🏼‍♀️",
	"🏄🏽‍♀️",
	"🏄🏾‍♀️",
	"🏄🏿‍♀️",
	"🏄‍♂️",
	"🏄🏻‍♂️",
	"🏄🏼‍♂️",
	"🏄🏽‍♂️",
	"🏄🏾‍♂️",
	"🏄🏿‍♂️",
	"🏊‍♀️",
	"🏊🏻‍♀️",
	"🏊🏼‍♀️",
	"🏊🏽‍♀️",
	"🏊🏾‍♀️",
	"🏊🏿‍♀️",
	"🏊‍♂️",
	"🏊🏻‍♂️",
	"🏊🏼‍♂️",
	"🏊🏽‍♂️",
	"🏊🏾‍♂️",
	"🏊🏿‍♂️",
	"🤽‍♀️",
	"🤽🏻‍♀️",
	"🤽🏼‍♀️",
	"🤽🏽‍♀️",
	"🤽🏾‍♀️",
	"🤽🏿‍♀️",
	"🤽‍♂️",
	"🤽🏻‍♂️",
	"🤽🏼‍♂️",
	"🤽🏽‍♂️",
	"🤽🏾‍♂️",
	"🤽🏿‍♂️",
	"🚣‍♀️",
	"🚣🏻‍♀️",
	"🚣🏼‍♀️",
	"🚣🏽‍♀️",
	"🚣🏾‍♀️",
	"🚣🏿‍♀️",
	"🚣‍♂️",
	"🚣🏻‍♂️",
	"🚣🏼‍♂️",
	"🚣🏽‍♂️",
	"🚣🏾‍♂️",
	"🚣🏿‍♂️",
	"🧗‍♀️",
	"🧗🏻‍♀️",
	"🧗🏼‍♀️",
	"🧗🏽‍♀️",
	"🧗🏾‍♀️",
	"🧗🏿‍♀️",
	"🧗‍♂️",
	"🧗🏻‍♂️",
	"🧗🏼‍♂️",
	"🧗🏽‍♂️",
	"🧗🏾‍♂️",
	"🧗🏿‍♂️",
	"🚵‍♀️",
	"🚵🏻‍♀️",
	"🚵🏼‍♀️",
	"🚵🏽‍♀️",
	"🚵🏾‍♀️",
	"🚵🏿‍♀️",
	"🚵‍♂️",
	"🚵🏻‍♂️",
	"🚵🏼‍♂️",
	"🚵🏽‍♂️",
	"🚵🏾‍♂️",
	"🚵🏿‍♂️",
	"🚴‍♀️",
	"🚴🏻‍♀️",
	"🚴🏼‍♀️",
	"🚴🏽‍♀️",
	"🚴🏾‍♀️",
	"🚴🏿‍♀️",
	"🚴‍♂️",
	"🚴🏻‍♂️",
	"🚴🏼‍♂️",
	"🚴🏽‍♂️",
	"🚴🏾‍♂️",
	"🚴🏿‍♂️",
	"🏆",
	"🥇",
	"🥈",
	"🥉",
	"🏅",
	"🎖",
	"🏵",
	"🎗",
	"🎫",
	"🎟",
	"🎪",
	"🤹‍♀️",
	"🤹🏻‍♀️",
	"🤹🏼‍♀️",
	"🤹🏽‍♀️",
	"🤹🏾‍♀️",
	"🤹🏿‍♀️",
	"🤹‍♂️",
	"🤹🏻‍♂️",
	"🤹🏼‍♂️",
	"🤹🏽‍♂️",
	"🤹🏾‍♂️",
	"🤹🏿‍♂️",
	"🎭",
	"🎨",
	"🎬",
	"🎤",
	"🎧",
	"🎼",
	"🎹",
	"🥁",
	"🎷",
	"🎺",
	"🎸",
	"🎻",
	"🎲",
	"🧩",
	"♟",
	"🎯",
	"🎳",
	"🎮",
	"🎰",
	"🚗",
	"🚕",
	"🚙",
	"🚌",
	"🚎",
	"🏎",
	"🚓",
	"🚑",
	"🚒",
	"🚐",
	"🚚",
	"🚛",
	"🚜",
	"🛴",
	"🚲",
	"🛵",
	"🏍",
	"🚨",
	"🚔",
	"🚍",
	"🚘",
	"🚖",
	"🚡",
	"🚠",
	"🚟",
	"🚃",
	"🚋",
	"🚞",
	"🚝",
	"🚄",
	"🚅",
	"🚈",
	"🚂",
	"🚆",
	"🚇",
	"🚊",
	"🚉",
	"✈️",
	"🛫",
	"🛬",
	"🛩",
	"💺",
	"🛰",
	"🚀",
	"🛸",
	"🚁",
	"🛶",
	"⛵️",
	"🚤",
	"🛥",
	"🛳",
	"⛴",
	"🚢",
	"⚓️",
	"⛽️",
	"🚧",
	"🚦",
	"🚥",
	"🚏",
	"🗺",
	"🗿",
	"🗽",
	"🗼",
	"🏰",
	"🏯",
	"🏟",
	"🎡",
	"🎢",
	"🎠",
	"⛲️",
	"⛱",
	"🏖",
	"🏝",
	"🏜",
	"🌋",
	"⛰",
	"🏔",
	"🗻",
	"🏕",
	"⛺️",
	"🏠",
	"🏡",
	"🏘",
	"🏚",
	"🏗",
	"🏭",
	"🏢",
	"🏬",
	"🏣",
	"🏤",
	"🏥",
	"🏦",
	"🏨",
	"🏪",
	"🏫",
	"🏩",
	"💒",
	"🏛",
	"⛪️",
	"🕌",
	"🕍",
	"🕋",
	"⛩",
	"🛤",
	"🛣",
	"🗾",
	"🎑",
	"🏞",
	"🌅",
	"🌄",
	"🌠",
	"🎇",
	"🎆",
	"🌇",
	"🌆",
	"🏙",
	"🌃",
	"🌌",
	"🌉",
	"🌁",
	"⌚️",
	"📱",
	"📲",
	"💻",
	"⌨️",
	"🖥",
	"🖨",
	"🖱",
	"🖲",
	"🕹",
	"🗜",
	"💽",
	"💾",
	"💿",
	"📀",
	"📼",
	"📷",
	"📸",
	"📹",
	"🎥",
	"📽",
	"🎞",
	"📞",
	"☎️",
	"📟",
	"📠",
	"📺",
	"📻",
	"🎙",
	"🎚",
	"🎛",
	"⏱",
	"⏲",
	"⏰",
	"🕰",
	"⌛️",
	"⏳",
	"📡",
	"🔋",
	"🔌",
	"💡",
	"🔦",
	"🕯",
	"🗑",
	"🛢",
	"💸",
	"💵",
	"💴",
	"💶",
	"💷",
	"💰",
	"💳",
	"🧾",
	"💎",
	"⚖️",
	"🔧",
	"🔨",
	"⚒",
	"🛠",
	"⛏",
	"🔩",
	"⚙️",
	"⛓",
	"🔫",
	"💣",
	"🔪",
	"🗡",
	"⚔️",
	"🛡",
	"🚬",
	"⚰️",
	"⚱️",
	"🏺",
	"🧭",
	"🧱",
	"🔮",
	"🧿",
	"🧸",
	"📿",
	"💈",
	"⚗️",
	"🔭",
	"🧰",
	"🧲",
	"🧪",
	"🧫",
	"🧬",
	"🧯",
	"🔬",
	"🕳",
	"💊",
	"💉",
	"🌡",
	"🚽",
	"🚰",
	"🚿",
	"🛁",
	"🛀",
	"🛀🏻",
	"🛀🏼",
	"🛀🏽",
	"🛀🏾",
	"🛀🏿",
	"🧴",
	"🧵",
	"🧶",
	"🧷",
	"🧹",
	"🧺",
	"🧻",
	"🧼",
	"🧽",
	"🛎",
	"🔑",
	"🗝",
	"🚪",
	"🛋",
	"🛏",
	"🛌",
	"🖼",
	"🛍",
	"🧳",
	"🛒",
	"🎁",
	"🎈",
	"🎏",
	"🎀",
	"🎊",
	"🎉",
	"🧨",
	"🎎",
	"🏮",
	"🎐",
	"🧧",
	"✉️",
	"📩",
	"📨",
	"📧",
	"💌",
	"📥",
	"📤",
	"📦",
	"🏷",
	"📪",
	"📫",
	"📬",
	"📭",
	"📮",
	"📯",
	"📜",
	"📃",
	"📄",
	"📑",
	"📊",
	"📈",
	"📉",
	"🗒",
	"🗓",
	"📆",
	"📅",
	"📇",
	"🗃",
	"🗳",
	"🗄",
	"📋",
	"📁",
	"📂",
	"🗂",
	"🗞",
	"📰",
	"📓",
	"📔",
	"📒",
	"📕",
	"📗",
	"📘",
	"📙",
	"📚",
	"📖",
	"🔖",
	"🔗",
	"📎",
	"🖇",
	"📐",
	"📏",
	"📌",
	"📍",
	"✂️",
	"🖊",
	"🖋",
	"✒️",
	"🖌",
	"🖍",
	"📝",
	"✏️",
	"🔍",
	"🔎",
	"🔏",
	"🔐",
	"🔒",
	"🔓",
	"❤️",
	"🧡",
	"💛",
	"💚",
	"💙",
	"💜",
	"🖤",
	"💔",
	"❣️",
	"💕",
	"💞",
	"💓",
	"💗",
	"💖",
	"💘",
	"💝",
	"💟",
	"☮️",
	"✝️",
	"☪️",
	"🕉",
	"☸️",
	"✡️",
	"🔯",
	"🕎",
	"☯️",
	"☦️",
	"🛐",
	"⛎",
	"♈️",
	"♉️",
	"♊️",
	"♋️",
	"♌️",
	"♍️",
	"♎️",
	"♏️",
	"♐️",
	"♑️",
	"♒️",
	"♓️",
	"🆔",
	"⚛️",
	"🉑",
	"☢️",
	"☣️",
	"📴",
	"📳",
	"🈶",
	"🈚️",
	"🈸",
	"🈺",
	"🈷️",
	"✴️",
	"🆚",
	"💮",
	"🉐",
	"㊙️",
	"㊗️",
	"🈴",
	"🈵",
	"🈹",
	"🈲",
	"🅰️",
	"🅱️",
	"🆎",
	"🆑",
	"🅾️",
	"🆘",
	"❌",
	"⭕️",
	"🛑",
	"⛔️",
	"📛",
	"🚫",
	"💯",
	"💢",
	"♨️",
	"🚷",
	"🚯",
	"🚳",
	"🚱",
	"🔞",
	"📵",
	"🚭",
	"❗️",
	"❕",
	"❓",
	"❔",
	"‼️",
	"⁉️",
	"🔅",
	"🔆",
	"〽️",
	"⚠️",
	"🚸",
	"🔱",
	"⚜️",
	"🔰",
	"♻️",
	"✅",
	"🈯️",
	"💹",
	"❇️",
	"✳️",
	"❎",
	"🌐",
	"💠",
	"Ⓜ️",
	"🌀",
	"💤",
	"🏧",
	"🚾",
	"♿️",
	"🅿️",
	"🈳",
	"🈂️",
	"🛂",
	"🛃",
	"🛄",
	"🛅",
	"🚹",
	"🚺",
	"🚼",
	"🚻",
	"🚮",
	"🎦",
	"📶",
	"🈁",
	"🔣",
	"ℹ️",
	"🔤",
	"🔡",
	"🔠",
	"🆖",
	"🆗",
	"🆙",
	"🆒",
	"🆕",
	"🆓",
	"0️⃣",
	"1️⃣",
	"2️⃣",
	"3️⃣",
	"4️⃣",
	"5️⃣",
	"6️⃣",
	"7️⃣",
	"8️⃣",
	"9️⃣",
	"🔟",
	"🔢",
	"#️⃣",
	"*️⃣",
	"⏏️",
	"▶️",
	"⏸",
	"⏯",
	"⏹",
	"⏺",
	"⏭",
	"⏮",
	"⏩",
	"⏪",
	"⏫",
	"⏬",
	"◀️",
	"🔼",
	"🔽",
	"➡️",
	"⬅️",
	"⬆️",
	"⬇️",
	"↗️",
	"↘️",
	"↙️",
	"↖️",
	"↕️",
	"↔️",
	"↪️",
	"↩️",
	"⤴️",
	"⤵️",
	"🔀",
	"🔁",
	"🔂",
	"🔄",
	"🔃",
	"🎵",
	"🎶",
	"➕",
	"➖",
	"➗",
	"✖️",
	"♾",
	"💲",
	"💱",
	"™️",
	"©️",
	"®️",
	"〰️",
	"➰",
	"➿",
	"🔚",
	"🔙",
	"🔛",
	"🔝",
	"🔜",
	"✔️",
	"☑️",
	"🔘",
	"⚪️",
	"⚫️",
	"🔴",
	"🔵",
	"🔺",
	"🔻",
	"🔸",
	"🔹",
	"🔶",
	"🔷",
	"🔳",
	"🔲",
	"▪️",
	"▫️",
	"◾️",
	"◽️",
	"◼️",
	"◻️",
	"⬛️",
	"⬜️",
	"🔈",
	"🔇",
	"🔉",
	"🔊",
	"🔔",
	"🔕",
	"📣",
	"📢",
	"👁‍🗨",
	"💬",
	"💭",
	"🗯",
	"♠️",
	"♣️",
	"♥️",
	"♦️",
	"🃏",
	"🎴",
	"🀄️",
	"🕐",
	"🕑",
	"🕒",
	"🕓",
	"🕔",
	"🕕",
	"🕖",
	"🕗",
	"🕘",
	"🕙",
	"🕚",
	"🕛",
	"🕜",
	"🕝",
	"🕞",
	"🕟",
	"🕠",
	"🕡",
	"🕢",
	"🕣",
	"🕤",
	"🕥",
	"🕦",
	"🕧",
	"🏳️",
	"🏴",
	"🏁",
	"🚩",
	"🏳️‍🌈",
	"🏴‍☠️",
	"🇦🇫",
	"🇦🇽",
	"🇦🇱",
	"🇩🇿",
	"🇦🇸",
	"🇦🇩",
	"🇦🇴",
	"🇦🇮",
	"🇦🇶",
	"🇦🇬",
	"🇦🇷",
	"🇦🇲",
	"🇦🇼",
	"🇦🇺",
	"🇦🇹",
	"🇦🇿",
	"🇧🇸",
	"🇧🇭",
	"🇧🇩",
	"🇧🇧",
	"🇧🇾",
	"🇧🇪",
	"🇧🇿",
	"🇧🇯",
	"🇧🇲",
	"🇧🇹",
	"🇧🇴",
	"🇧🇦",
	"🇧🇼",
	"🇧🇷",
	"🇮🇴",
	"🇻🇬",
	"🇧🇳",
	"🇧🇬",
	"🇧🇫",
	"🇧🇮",
	"🇰🇭",
	"🇨🇲",
	"🇨🇦",
	"🇮🇨",
	"🇨🇻",
	"🇧🇶",
	"🇰🇾",
	"🇨🇫",
	"🇹🇩",
	"🇨🇱",
	"🇨🇳",
	"🇨🇽",
	"🇨🇨",
	"🇨🇴",
	"🇰🇲",
	"🇨🇬",
	"🇨🇩",
	"🇨🇰",
	"🇨🇷",
	"🇨🇮",
	"🇭🇷",
	"🇨🇺",
	"🇨🇼",
	"🇨🇾",
	"🇨🇿",
	"🇩🇰",
	"🇩🇯",
	"🇩🇲",
	"🇩🇴",
	"🇪🇨",
	"🇪🇬",
	"🇸🇻",
	"🇬🇶",
	"🇪🇷",
	"🇪🇪",
	"🇪🇹",
	"🇪🇺",
	"🇫🇰",
	"🇫🇴",
	"🇫🇯",
	"🇫🇮",
	"🇫🇷",
	"🇬🇫",
	"🇵🇫",
	"🇹🇫",
	"🇬🇦",
	"🇬🇲",
	"🇬🇪",
	"🇩🇪",
	"🇬🇭",
	"🇬🇮",
	"🇬🇷",
	"🇬🇱",
	"🇬🇩",
	"🇬🇵",
	"🇬🇺",
	"🇬🇹",
	"🇬🇬",
	"🇬🇳",
	"🇬🇼",
	"🇬🇾",
	"🇭🇹",
	"🇭🇳",
	"🇭🇰",
	"🇭🇺",
	"🇮🇸",
	"🇮🇳",
	"🇮🇩",
	"🇮🇷",
	"🇮🇶",
	"🇮🇪",
	"🇮🇲",
	"🇮🇱",
	"🇮🇹",
	"🇯🇲",
	"🇯🇵",
	"🎌",
	"🇯🇪",
	"🇯🇴",
	"🇰🇿",
	"🇰🇪",
	"🇰🇮",
	"🇽🇰",
	"🇰🇼",
	"🇰🇬",
	"🇱🇦",
	"🇱🇻",
	"🇱🇧",
	"🇱🇸",
	"🇱🇷",
	"🇱🇾",
	"🇱🇮",
	"🇱🇹",
	"🇱🇺",
	"🇲🇴",
	"🇲🇰",
	"🇲🇬",
	"🇲🇼",
	"🇲🇾",
	"🇲🇻",
	"🇲🇱",
	"🇲🇹",
	"🇲🇭",
	"🇲🇶",
	"🇲🇷",
	"🇲🇺",
	"🇾🇹",
	"🇲🇽",
	"🇫🇲",
	"🇲🇩",
	"🇲🇨",
	"🇲🇳",
	"🇲🇪",
	"🇲🇸",
	"🇲🇦",
	"🇲🇿",
	"🇲🇲",
	"🇳🇦",
	"🇳🇷",
	"🇳🇵",
	"🇳🇱",
	"🇳🇨",
	"🇳🇿",
	"🇳🇮",
	"🇳🇪",
	"🇳🇬",
	"🇳🇺",
	"🇳🇫",
	"🇰🇵",
	"🇲🇵",
	"🇳🇴",
	"🇴🇲",
	"🇵🇰",
	"🇵🇼",
	"🇵🇸",
	"🇵🇦",
	"🇵🇬",
	"🇵🇾",
	"🇵🇪",
	"🇵🇭",
	"🇵🇳",
	"🇵🇱",
	"🇵🇹",
	"🇵🇷",
	"🇶🇦",
	"🇷🇪",
	"🇷🇴",
	"🇷🇺",
	"🇷🇼",
	"🇼🇸",
	"🇸🇲",
	"🇸🇦",
	"🇸🇳",
	"🇷🇸",
	"🇸🇨",
	"🇸🇱",
	"🇸🇬",
	"🇸🇽",
	"🇸🇰",
	"🇸🇮",
	"🇬🇸",
	"🇸🇧",
	"🇸🇴",
	"🇿🇦",
	"🇰🇷",
	"🇸🇸",
	"🇪🇸",
	"🇱🇰",
	"🇧🇱",
	"🇸🇭",
	"🇰🇳",
	"🇱🇨",
	"🇵🇲",
	"🇻🇨",
	"🇸🇩",
	"🇸🇷",
	"🇸🇿",
	"🇸🇪",
	"🇨🇭",
	"🇸🇾",
	"🇹🇼",
	"🇹🇯",
	"🇹🇿",
	"🇹🇭",
	"🇹🇱",
	"🇹🇬",
	"🇹🇰",
	"🇹🇴",
	"🇹🇹",
	"🇹🇳",
	"🇹🇷",
	"🇹🇲",
	"🇹🇨",
	"🇹🇻",
	"🇻🇮",
	"🇺🇬",
	"🇺🇦",
	"🇦🇪",
	"🇬🇧",
	"🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	"🏴󠁧󠁢󠁳󠁣󠁴󠁿",
	"🏴󠁧󠁢󠁷󠁬󠁳󠁿",
	"🇺🇳",
	"🇺🇸",
	"🇺🇾",
	"🇺🇿",
	"🇻🇺",
	"🇻🇦",
	"🇻🇪",
	"🇻🇳",
	"🇼🇫",
	"🇪🇭",
	"🇾🇪",
	"🇿🇲",
	"🇿🇼",
	"🥰",
	"🥵",
	"🥶",
	"🥳",
	"🥴",
	"🥺",
	"👨‍🦰",
	"👩‍🦰",
	"👨‍🦱",
	"👩‍🦱",
	"👨‍🦲",
	"👩‍🦲",
	"👨‍🦳",
	"👩‍🦳",
	"🦸",
	"🦸‍♀️",
	"🦸‍♂️",
	"🦹",
	"🦹‍♀️",
	"🦹‍♂️",
	"🦵",
	"🦶",
	"🦴",
	"🦷",
	"🥽",
	"🥼",
	"🥾",
	"🥿",
	"🦝",
	"🦙",
	"🦛",
	"🦘",
	"🦡",
	"🦢",
	"🦚",
	"🦜",
	"🦞",
	"🦟",
	"🦠",
	"🥭",
	"🥬",
	"🥯",
	"🧂",
	"🥮",
	"🧁",
	"🧭",
	"🧱",
	"🛹",
	"🧳",
	"🧨",
	"🧧",
	"🥎",
	"🥏",
	"🥍",
	"🧿",
	"🧩",
	"🧸",
	"♟",
	"🧮",
	"🧾",
	"🧰",
	"🧲",
	"🧪",
	"🧫",
	"🧬",
	"🧯",
	"🧴",
	"🧵",
	"🧶",
	"🧷",
	"🧹",
	"🧺",
	"🧻",
	"🧼",
	"🧽",
	"♾",
	"🏴‍☠️",
}

func CountEmoji(src string) (m map[string]int) {
	m = make(map[string]int)
	for _, v := range emojis {
		c := strings.Count(src, v)
		if c > 0 {
			m[v] = c
		}
	}

	return
}
