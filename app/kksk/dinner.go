package main

import "math/rand"
import . "github.com/maxence-charriere/go-app/v9/pkg/app"

type dinner struct {
	Compo

	fallacy rhetologicalFallacy
}

func newDinner() *dinner {
	return &dinner{}
}

func (c *dinner) Render() UI {
	if c.fallacy.Image == "" {
		c.fallacy = getRandomRhetologicalFallacy()
	}
	return Div().Body(
		Div().Body(
			Img().Src("/web/Rhetological-Fallacies/"+c.fallacy.Image).
				Styles(map[string]string{
					"height": "230px",
					"width":  "315px",
				}),
		),

		Div().Body(
			H2().Text(c.fallacy.Title).
				Styles(map[string]string{
					"height":    "25px",
					"font-size": "27px",
				}),
			Div().
				Body(
					P().Text(c.fallacy.Description).Styles(map[string]string{
						"font-size": "16px",
					}),
					P().Text(c.fallacy.ShadyWords).Styles(map[string]string{
						"font-size": "13px",
					}),
				).Styles(map[string]string{
				"height": "86px",
			}),
		).Styles(map[string]string{
			"height": "230px",
			"width":  "315px",
		}),

		Div().Body(
			Button().Text("换一张").
				OnClick(func(ctx Context, e Event) {
					c.fallacy = getRandomRhetologicalFallacy()
				}),
		),
	)
}

type rhetologicalFallacy struct {
	Image       string
	Title       string
	Description string
	ShadyWords  string
}

func getRandomRhetologicalFallacy() rhetologicalFallacy {
	random := rand.Intn(len(rhetologicalFallacies))
	for k := range rhetologicalFallacies {
		if random == 0 {
			return rhetologicalFallacies[k]
		}
		random--
	}
	return rhetologicalFallacy{}
}

var rhetologicalFallacies = map[string]rhetologicalFallacy{
	"诡辩术": {
		Image:       "Rhetological-Fallacies-00-HEADER-CH.png",
		Title:       "这里写点垃圾话",
		Description: "这里解释一下垃圾话",
		ShadyWords:  "“最后阴阳怪气一句，嘻嘻”",
	},
	"掩耳盗铃": {
		Image:       "Rhetological-Fallacies-01-Belief.png",
		Title:       "掩耳盗铃",
		Description: "声称某个观点是错误的，因为你不愿意相信那所意味的事实",
		ShadyWords:  "“他不可能就是为了骗我的钱。他说过他爱我的，他一定是遭到了什么变故。”",
	},
	"诉诸恐惧": {
		Image:       "Rhetological-Fallacies-02-Fear.png",
		Title:       "诉诸恐惧",
		Description: "煽动对一方的恐惧与偏见，从而进行论证。",
		ShadyWords:  "“在你意识到之前，清真寺的数目就会超过教堂了。”",
	},
	"诉诸谄媚": {
		Image:       "Rhetological-Fallacies-03-Flattery.png",
		Title:       "诉诸谄媚",
		Description: "给毫无根据的论点裹上糖衣炮弹，让人不自觉地全盘接受。",
		ShadyWords:  "“聪明又有洞察力的读者当然在读到这样的谬误时就能马上发觉。”",
	},
	"诉诸自然": {
		Image:       "Rhetological-Fallacies-04-Nature.png",
		Title:       "诉诸自然",
		Description: "通过与“至善”的自然界的对比，来使你的观点看起来更站得住脚。",
		ShadyWords:  "“同性恋当然违背天性。你看不到同性动物交配吧。”",
	},
	"诉诸同情": {
		Image:       "Rhetological-Fallacies-05-Pity.png",
		Title:       "诉诸同情",
		Description: "唤起人们的怜悯之心，以动摇对手。",
		ShadyWords:  "“前独裁者已垂垂老矣，濒临末年。让他为这些指控接受审判实在不应该。”",
	},
	"诉诸荒谬": {
		Image:       "Rhetological-Fallacies-06-Ridicule.png",
		Title:       "诉诸荒谬",
		Description: "将对手的观点以荒谬的形式表现出来以进行打击。",
		ShadyWords:  "“对上帝的忠诚就如同相信有圣诞老人和牙仙一样。”",
	},
	"诉诸仇恨": {
		Image:       "Rhetological-Fallacies-07-Spite.png",
		Title:       "诉诸仇恨",
		Description: "因为个人偏见而对某一看法不屑一顾。",
		ShadyWords:  "“富二代搞慈善？算了吧，反正还不是在做秀。”",
	},
	"一厢情愿": {
		Image:       "Rhetological-Fallacies-08-Wishful.png",
		Title:       "一厢情愿",
		Description: "认为一件事是真的或假的，仅仅因为你情愿想当然。",
		ShadyWords:  "“主席是不会犯错的。他是人民的领袖，红旗的舵手。”",
	},

	"诉诸匿名权威": {
		Image:       "Rhetological-Fallacies-09-Anonymous.png",
		Title:       "诉诸匿名权威",
		Description: "引用来源不详的“砖家”、“研究”或某一群体（比如“科学家”）以声称某观点是正确的。",
		ShadyWords:  "“他们说要花7年才能消化一片口香糖。”",
	},
	"诉诸（可疑）权威": {
		Image:       "Rhetological-Fallacies-10-Authority.png",
		Title:       "诉诸（可疑）权威",
		Description: "因为某个无实无信的“专家”说某件事是真的，因而断言确实如此。",
		ShadyWords:  "“超过400位杰出的科学家与工程师对全球变暖持争议态度。”",
	},
	"诉诸常规": {
		Image:       "Rhetological-Fallacies-11-CommonPractice.png",
		Title:       "诉诸常规",
		Description: "因为常见，所以正确。",
		ShadyWords:  "“这家银行有些贪污腐败方面的问题。但在这里发生的事不是在哪家银行都在发生嘛。”",
	},
	"诉诸无知": {
		Image:       "Rhetological-Fallacies-12-Ignorance.png",
		Title:       "诉诸无知",
		Description: "某一观点是正确的，仅仅因为它没有被证伪（或某观点是错误的，仅仅因为它尚未被证实）",
		ShadyWords:  "“没人能证明有上帝。所以没有上帝。”",
	},
	"诉诸怀疑": {
		Image:       "Rhetological-Fallacies-13-Incredulity.png",
		Title:       "诉诸怀疑",
		Description: "因为某件事听起来不可信，所以一定不是真的。",
		ShadyWords:  "“眼睛是超级复杂的生物机械之作，有千万个紧密联系的部件。如果没有一位睿智的设计师，这怎么可能存在？”",
	},
	"身价逻辑": {
		Image:       "Rhetological-Fallacies-14-Money.png",
		Title:       "身价逻辑",
		Description: "如果某人很有钱，或者某样东西很贵，那么这就对某一论断的真实性造成了影响。",
		ShadyWords:  "“如果这玩意儿更贵的话，那它一定更好。”",
	},
	"求新逻辑": {
		Image:       "Rhetological-Fallacies-15-Novelty.png",
		Title:       "求新逻辑",
		Description: "因为是最新的，所以更好。",
		ShadyWords:  "“太棒了！最新的操作系统会让我的电脑跑得更快的…”",
	},
	"诉诸主流": {
		Image:       "Rhetological-Fallacies-16-Popular.png",
		Title:       "诉诸主流",
		Description: "认定某件事是真的，因为大多数人都这么相信。",
		ShadyWords:  "“喝牛奶能使你骨骼强健。”",
	},
	"诉诸概率": {
		Image:       "Rhetological-Fallacies-17-Probability.png",
		Title:       "诉诸概率",
		Description: "相信因为某件事情可能发生，所以必然会发生。",
		ShadyWords:  "“宇宙里有数不清的星系，无数多的星星。一定有另一颗行星孕育了智慧生命。”",
	},
	"诉诸传统": {
		Image:       "Rhetological-Fallacies-18-Tradition.png",
		Title:       "诉诸传统",
		Description: "声称某件事是正确的，因为（很显然）一直以来都是这样。",
		ShadyWords:  "“婚姻是男与女的结合。因此同性婚姻毫无根据。”",
	},
	"轶事证据": {
		Image:       "Rhetological-Fallacies-19-Anecdotal.png",
		Title:       "轶事证据",
		Description: "对系统性研究下得出的证据视而不见，反而集中在手头的个例上。",
		ShadyWords:  "“我才不戒烟呢。我爷爷每天抽40根，还活到了90岁！”",
	},
	"合成谬误": {
		Image:       "Rhetological-Fallacies-20-Composition.png",
		Title:       "合成谬误",
		Description: "推断一群人的特性或信条也代表了整个团体。",
		ShadyWords:  "“最近的恐怖袭击是由激进的伊斯兰教徒组织的。因此所有的恐怖分子都是穆斯林。”",
	},
	"分割谬误": {
		Image:       "Rhetological-Fallacies-21-Division.png",
		Title:       "分割谬误",
		Description: "将整个团体的特性或信条自动代入到每一名成员的头上。",
		ShadyWords:  "“苹果的产品向来颠覆传统，设计一流。下一款也一定如此。”",
	},
	"设计谬误": {
		Image:       "Rhetological-Fallacies-22-Design.png",
		Title:       "设计谬误",
		Description: "因为某样东西设计精美，视效上佳，所以更加站得住脚。",
		ShadyWords:  "“呃…”",
	},
	"赌徒谬误": {
		Image:       "Rhetological-Fallacies-23-Gambler.png",
		Title:       "赌徒谬误",
		Description: "认为历史结果会影响未来结果。",
		ShadyWords:  "“我已连续丢了10次硬币，都是正面朝上。因此下一次更可能丢出反面来。”",
	},
	"轻率概化": {
		Image:       "Rhetological-Fallacies-24-Hasty.png",
		Title:       "轻率概化",
		Description: "从单一的样本得出概括性的结论。",
		ShadyWords:  "“我被前面的女驾驶别了下。女人就是不能开车。”",
	},
	"妄下定论": {
		Image:       "Rhetological-Fallacies-25-Jumping.png",
		Title:       "妄下定论",
		Description: "没有公平考虑所有相关（且易举证的）事实，就妄下结论",
		ShadyWords:  "“她想要医疗保险报销避孕药？真是个婊子。”",
	},
	"中间立场": {
		Image:       "Rhetological-Fallacies-26-Middle.png",
		Title:       "中间立场",
		Description: "相冲突的两个观点似乎都有道理，那么答案一定在两者的中间地带。",
		ShadyWords:  "“我追尾了你的车，但我不认为自己该出修理费。你认为我该出所有的修理费。合乎情理的方案就是平分费用。”",
	},
	"完美主义谬误": {
		Image:       "Rhetological-Fallacies-27-Perfectionist.png",
		Title:       "完美主义谬误",
		Description: "认为只有完美的成功才是可行的选择，从而反对任何低于预期的方案。",
		ShadyWords:  "“这反酒驾的宣传究竟有什么用？人们还是会醉酒驾车的。”",
	},
	"相对论谬误": {
		Image:       "Rhetological-Fallacies-28-Relativist.png",
		Title:       "相对论谬误",
		Description: "否定某样客观事实，认为事实是相对一个或一群人而言的。",
		ShadyWords:  "“那对你来说可能是对的。但对我来说不是。”",
	},
	"以偏概全": {
		Image:       "Rhetological-Fallacies-29-Spotlight.png",
		Title:       "以偏概全",
		Description: "认为从小样本观察到的同样适用于整体。",
		ShadyWords:  "“这家大型制鞋商在血汗工厂里雇用童工。可想而知所有制鞋公司都是邪恶的童工奴隶主！”",
	},
	"一概而论": {
		Image:       "Rhetological-Fallacies-30-Sweeping.png",
		Title:       "一概而论",
		Description: "宽泛地应用一般性原则。",
		ShadyWords:  "“那些年轻人暴乱是因为他们缺失有道德观念的父亲。”",
	},
	"中词不周延": {
		Image:       "Rhetological-Fallacies-31-Undistributed.png",
		Title:       "中词不周延",
		Description: "因两件事有一个共通点，那么他们就是同一回事。",
		ShadyWords:  "“理论是尚未证实的观点。科学家用‘进化论’这一词，可见进化是未被证实的。”",
	},
	"临阵救援": {
		Image:       "Rhetological-Fallacies-32-Adhoc.png",
		Title:       "临阵救援",
		Description: "通过不断修改论据，搪塞问题，来保全自己的一贯主张。",
		ShadyWords:  "“…但除了更好的卫生，医药，教育，灌溉，公共卫生，道路，净水系统和公共秩序…罗马人为我们做了什么？”",
	},
	"一孔之见": {
		Image:       "Rhetological-Fallacies-33-Biased.png",
		Title:       "一孔之见",
		Description: "用不具代表性的样本所得出的结论，来支持自己的论点。",
		ShadyWords:  "“我们的网上调查表示，90%的互联网用户反对网络隐私法。”",
	},
	"确认偏误": {
		Image:       "Rhetological-Fallacies-34-Confirmation.png",
		Title:       "确认偏误",
		Description: "挑拣对自己有利的证据，而故意无视相冲突的。",
		ShadyWords:  "“很明显911事件是美国政府为了合理化伊拉克与阿富汗战争而发动的阴谋。没有飞机撞上五角大楼。双子塔的倒塌是控制爆破。。。等等”",
	},
	"伪二分法": {
		Image:       "Rhetological-Fallacies-35-Dilemma.png",
		Title:       "伪二分法",
		Description: "隐藏其它可能性，将两个对立的观点看作仅有的选择。",
		ShadyWords:  "“我们要么就得削减教育预算，要么就得负更多的债。我们不能负更多的债了，所以我们非降低教育预算不可。”",
	},
	"谎言": {
		Image:       "Rhetological-Fallacies-36-Lie.png",
		Title:       "谎言",
		Description: "彻头彻尾的谎言，被作为真相一再重复。",
		ShadyWords:  "“我没有和那个女人发生性关系。”",
	},
	"误导性鲜活个案": {
		Image:       "Rhetological-Fallacies-37-Vividness.png",
		Title:       "误导性鲜活个案",
		Description: "用生动的细节来描述一次小概率事件，以让别人相信这是一个问题。",
		ShadyWords:  "“在法院判决同性婚姻合法化之后，学校图书馆被要求存有同性文学作品；小学生会读到同性恋的童话故事，甚至有明确支持同性恋的手册。”",
	},
	"转移注意": {
		Image:       "Rhetological-Fallacies-38-Redherring.png",
		Title:       "转移注意",
		Description: "将毫不相关的话题引入辩论，以干扰视线并导向不同的结论。",
		ShadyWords:  "“参议员不需要为他开销的异常做出说明。毕竟，有些参议员做的破事儿比这严重多了。”",
	},
	"滑坡谬误": {
		Image:       "Rhetological-Fallacies-39-Slippery.png",
		Title:       "滑坡谬误",
		Description: "认为开始的一小步会无可避免地引发一串相关（负面）的事件。",
		ShadyWords:  "“如果我们将大麻合法化，更多的人就会开始吸食毒品和海洛因。到时候我们就得也合法化那些。”",
	},
	"隐瞒证据": {
		Image:       "Rhetological-Fallacies-40-Suppressed.png",
		Title:       "隐瞒证据",
		Description: "有意不用相关且重要的信息，因为那对立于自己的结论。",
		ShadyWords:  "“炒菜产生的油烟是PM2.5的重要来源，所以要治理雾霾，中国人就得少做菜。”",
	},
	"无法证伪": {
		Image:       "Rhetological-Fallacies-41-Unfalsiability.png",
		Title:       "无法证伪",
		Description: "提出一个无法被证伪的观点，因为无法加以验证。",
		ShadyWords:  "“他撒谎是因为鬼上身了。”",
	},
	"肯定后件": {
		Image:       "Rhetological-Fallacies-42-Consequent.png",
		Title:       "肯定后件",
		Description: "认为对你所观察到的现象只有一种解释。",
		ShadyWords:  "“婚姻会带来孩子的降生。所以这就是其存在的理由。”",
	},
	"循环逻辑": {
		Image:       "Rhetological-Fallacies-43-Circular.png",
		Title:       "循环逻辑",
		Description: "论证的前提里已经蕴含结论。",
		ShadyWords:  "“《圣经》上说上帝存在。由于圣经是上帝的话，圣经必然正确。所以上帝是存在的。”",
	},
	"相关即因果": {
		Image:       "Rhetological-Fallacies-44-CumHoc.png",
		Title:       "相关即因果",
		Description: "认为两个一起发生的事件一定有因果关系。（关联性=因果性）",
		ShadyWords:  "“小混混们听主题暴力的饶舌音乐。所以饶舌音乐会造成青少年的暴力行为。”",
	},
	"否定前件": {
		Image:       "Rhetological-Fallacies-45-Antecedent.png",
		Title:       "否定前件",
		Description: "有这样的结果并非只有一个解释。因此，在这样的情况下从结果反推原因是不准确的。",
		ShadyWords:  "“如果你读了好学校，你就会找到好工作。如果你没读好学校，你就找不到好工作。”",
	},
	"忽视主因": {
		Image:       "Rhetological-Fallacies-46-Ignoring.png",
		Title:       "忽视主因",
		Description: "声称是某事件导致了后果，而实际上另一件（意料之外 ）的事才是原因。",
		ShadyWords:  "“60年代我们开始了性解放运动，而现在人们正死于艾滋。”",
	},
	"前后即因果": {
		Image:       "Rhetological-Fallacies-47-PostHoc.png",
		Title:       "前后即因果",
		Description: "因为一件事是在另一件事之后发生的，因此也是由那件事引起的。",
		ShadyWords:  "“总统上台之后，失业人数创了历史新高。所以总统阻碍了经济发展。”",
	},
	"积非成是": {
		Image:       "Rhetological-Fallacies-48-TwoWrongs.png",
		Title:       "积非成是",
		Description: "认为一桩错事能被另一桩错事所抵消。",
		ShadyWords:  "“不错——这监狱环境恶劣又没人性，不过关的本来就是罪犯！”",
	},
	"人身攻击": {
		Image:       "Rhetological-Fallacies-49-AdHominem.png",
		Title:       "人身攻击",
		Description: "绕开论证，针对辩论者本身发起不相干的攻击。",
		ShadyWords:  "“你以为自己是生物学专家吗，也好意思来教我们转基因食品的事？”",
	},
	"举证责任": {
		Image:       "Rhetological-Fallacies-50-Burden.png",
		Title:       "举证责任",
		Description: "我不需要证明我说的正确——你必须证明它是错的。",
		ShadyWords:  "“我坚持认为长期的太阳活动周期是全球变暖的原因。证明我错了啊。”",
	},
	"身份主观": {
		Image:       "Rhetological-Fallacies-51-Circumstance.png",
		Title:       "身份主观",
		Description: "认为一个论断不可信，因为支持者与之有利益关系。",
		ShadyWords:  "“研究手机对健康影响的这个调研有手机公司参与。所以，研究结果不可信。”",
	},
	"基因谬误": {
		Image:       "Rhetological-Fallacies-52-Genetic.png",
		Title:       "基因谬误",
		Description: "攻击一个论点的来源，而非它的内容。",
		ShadyWords:  "“这本书是1967年出版的，里面说的东西哪还有价值。”",
	},
	"罪恶关联": {
		Image:       "Rhetological-Fallacies-53-Guilt.png",
		Title:       "罪恶关联",
		Description: "通过将一个论点与形象不良的人或群体联系起来，从而破坏其可信度。",
		ShadyWords:  "“哦，你想要放松反恐条例，就像那帮恐怖分子想要的一样。所以你是支持恐怖主义的啰？”",
	},
	"稻草人谬误": {
		Image:       "Rhetological-Fallacies-54-StrawMan.png",
		Title:       "稻草人谬误",
		Description: "歪曲或简化你对手的论点，以攻击之。",
		ShadyWords:  "甲：“国家应该投入更多的预算来发展教育行业。”\n                                                        乙：“你这么不爱国，居然想减少国防开支，让外国列强有机可乘。”",
	},
}
