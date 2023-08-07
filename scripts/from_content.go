package main

import (
	"github.com/gosimple/slug"
	"github.com/ozbeksu/samarkand-api/ent"
	"strings"
)

var (
	commentList   []map[string]any
	communityList []map[string]any
	tagList       []map[string]string
	topicList     []map[string]string
)

func createCommentWith(m map[string]any, userCount, imgCount int) *ent.CommentCreate {
	d := getRandDate()
	hs, bs := getScores(d)

	uID := getRandIntInRange(1, userCount)
	comment := db.Comment.Create().
		SetHotScore(hs).
		SetBestScore(bs).
		SetCreatedAt(d).
		SetAuthorID(uID)

	cID, ok := m["community_id"].(int)
	if ok {
		comment = comment.SetCommunityID(cID)
	}

	tagIDs, ok := m["tags"].([]int)
	if ok && (tagIDs != nil || len(tagIDs) > 0) {
		comment = comment.AddTagIDs(tagIDs...)
	}

	var co *ent.Content
	if m["parent_id"] == nil {
		t, ok := m["title"].(string)
		if ok && len(t) > 0 {
			s := slug.Make(strings.ToLower(t))
			comment = comment.SetTitle(t).SetSlug(s)
		}

		b := m["content"].(string)
		if getRandBool() {
			aID := getRandIntInRange(1, imgCount)
			co = createPostContentWithAttachment(b, aID)
			comment = comment.AddAttachmentIDs(aID)
			db.User.UpdateOneID(uID).AddAttachmentIDs(aID).SaveX(ctx)
		} else {
			co = createPostContent(b)
		}
		comment = comment.SetContentID(co.ID)
	} else {
		pID := m["parent_id"].(int)
		b := m["content"].(string)
		co = createPostContent(b)
		comment = comment.SetContentID(co.ID).SetParentID(pID)
	}

	return comment
}

func makeCommentsFromContent(userCount, imgCount int) []*ent.Comment {
	bulk := make([]*ent.CommentCreate, len(commentList))
	for i := 0; i < len(commentList); i++ {
		bulk[i] = createCommentWith(commentList[i], userCount, imgCount)
	}
	return db.Comment.CreateBulk(bulk...).SaveX(ctx)
}

func createCommunityWith(m map[string]any, topicCount, memberCount int) *ent.CommunityCreate {
	t := m["title"].(string)
	s := m["slug"].(string)
	d := m["description"].(string)

	creatorID := getRandIntInRange(1, memberCount)
	topicIDs := []int{getRandIntInRange(1, topicCount), getRandIntInRange(1, topicCount)}
	memberIDs := []int{getRandIntInRange(1, memberCount), getRandIntInRange(1, memberCount), getRandIntInRange(1, memberCount)}
	modIDs := []int{getRandIntInRange(1, memberCount), getRandIntInRange(1, memberCount)}

	return db.Community.Create().
		SetTitle(t).
		SetCreatorID(creatorID).
		SetSlug(slug.Make(s)).
		SetDescription(d).
		AddTopicIDs(topicIDs...).
		AddMemberIDs(memberIDs...).
		AddModeratorIDs(modIDs...)
}

func makeCommunitiesFromContent(topicCount, memberCount int) []*ent.Community {
	bulk := make([]*ent.CommunityCreate, len(communityList))
	for i := 0; i < len(communityList); i++ {
		bulk[i] = createCommunityWith(communityList[i], topicCount, memberCount)
	}
	return db.Community.CreateBulk(bulk...).SaveX(ctx)
}

func createTagWith(m map[string]string) *ent.TagCreate {
	s := slug.Make(strings.ToLower(m["name"]))

	return db.Tag.Create().SetName(m["name"]).SetSlug(s)
}

func makeTagsFromContent() []*ent.Tag {
	bulk := make([]*ent.TagCreate, len(tagList))
	for i := 0; i < len(tagList); i++ {
		bulk[i] = createTagWith(tagList[i])
	}
	return db.Tag.CreateBulk(bulk...).SaveX(ctx)
}

func createTopicWith(m map[string]string) *ent.TopicCreate {
	s := slug.Make(strings.ToLower(m["name"]))

	return db.Topic.Create().SetName(m["name"]).SetSlug(s)
}

func makeTopicsFromContent() []*ent.Topic {
	bulk := make([]*ent.TopicCreate, len(topicList))
	for i := 0; i < len(topicList); i++ {
		bulk[i] = createTopicWith(topicList[i])
	}
	return db.Topic.CreateBulk(bulk...).SaveX(ctx)
}

func init() {
	commentList = []map[string]any{
		{
			"id":           1,
			"parent_id":    nil,
			"user_id":      1,
			"community_id": 1,
			"title":        "I just saw the latest demo of the XYZ Virtual Reality headset.",
			"content":      "The level of immersion is unlike anything I've ever seen before! Has anyone tried it?",
			"tags":         []int{1, 2},
			"comments":     []int{2},
		},
		{
			"id":           2,
			"user_id":      2,
			"parent_id":    1,
			"community_id": 1,
			"content":      "I've tried it at a tech expo, and I totally agree! The motion tracking is spot-on. Have you seen the accessories they're launching with it?",
			"tags":         nil,
			"comments":     []int{3},
		},
		{
			"id":           3,
			"user_id":      3,
			"parent_id":    2,
			"community_id": 1,
			"content":      "I heard there's a haptic feedback glove coming with it. I can only imagine what gaming would feel like with that level of immersion.",
			"tags":         nil,
			"comments":     []int{4},
		},
		{
			"id":           4,
			"user_id":      1,
			"parent_id":    3,
			"community_id": 1,
			"content":      "Yes! The glove is part of the bundle. It's truly a game-changer for the VR industry. Can't wait to see what developers do with this technology.",
			"tags":         nil,
			"comments":     []int{5},
		},
		{
			"id":           5,
			"user_id":      4,
			"parent_id":    4,
			"community_id": 1,
			"content":      "I'm curious about the price point and whether it will be accessible to the average consumer. Any idea on the pricing?",
			"tags":         nil,
			"comments":     []int{6},
		},
		{
			"id":           6,
			"user_id":      2,
			"parent_id":    5,
			"community_id": 1,
			"content":      "I believe they're aiming for a competitive price, especially considering the advanced features. It might still be on the higher end, but definitely something to watch for during sales!",
			"tags":         nil,
			"comments":     nil,
		},
		{
			"id":           7,
			"user_id":      5,
			"parent_id":    nil,
			"community_id": 1,
			"title":        "Has anyone tried the new ABC Smart Home System?",
			"content":      "I'm considering it for my house but would love to hear some experiences first.",
			"tags":         []int{1, 2},
			"comments":     []int{8},
		},
		{
			"id":           8,
			"user_id":      6,
			"parent_id":    7,
			"community_id": 1,
			"content":      "I've installed it last month, and it's incredible. The automation and voice control features are top-notch. Definitely worth considering.",
			"tags":         nil,
			"comments":     []int{9},
		},
		{
			"id":           9,
			"user_id":      5,
			"parent_id":    8,
			"community_id": 1,
			"content":      "That's great to hear! How was the setup process? Did you encounter any issues?",
			"tags":         nil,
			"comments":     []int{10},
		},
		{
			"id":           10,
			"user_id":      6,
			"parent_id":    9,
			"community_id": 1,
			"content":      "The setup was fairly straightforward. The app guides you through it step by step. Just make sure to follow the instructions carefully.",
			"tags":         nil,
			"comments":     []int{11},
		},
		{
			"id":           11,
			"user_id":      7,
			"parent_id":    10,
			"community_id": 1,
			"content":      "I'm considering the ABC Smart Home System too. How is their customer support? I heard mixed reviews about it.",
			"tags":         nil,
			"comments":     []int{12},
		},
		{
			"id":           12,
			"user_id":      6,
			"parent_id":    11,
			"community_id": 1,
			"content":      "Customer support was helpful when I had a question. They have chat support on their website, which I found to be quite efficient.",
			"tags":         nil,
			"comments":     nil,
		},
		{
			"id":           13,
			"user_id":      8,
			"parent_id":    nil,
			"community_id": 1,
			"title":        "I came across a new startup working on quantum computing.",
			"content":      "They claim to have a working prototype that can solve complex problems in seconds. Has anyone else heard about this? #GadgetGuru #FutureTech",
			"tags":         []int{2, 3},
			"comments":     []int{14},
		},
		{
			"id":           14,
			"user_id":      9,
			"parent_id":    13,
			"community_id": 1,
			"content":      "Yes, I read about it too! Quantum computing is a fascinating field with immense potential. Do you know when they are planning to reveal it to the public?",
			"tags":         nil,
			"comments":     []int{15},
		},
		{
			"id":           15,
			"user_id":      8,
			"parent_id":    14,
			"community_id": 1,
			"content":      "They are planning a big reveal next month. I'm eagerly waiting to see what they have accomplished. It could revolutionize many industries.",
			"tags":         nil,
			"comments":     []int{16},
		},
		{
			"id":           16,
			"user_id":      10,
			"parent_id":    15,
			"community_id": 1,
			"content":      "Quantum computing is still in its infancy, but the progress made by startups is encouraging. I hope it becomes more accessible to researchers and scientists soon.",
			"tags":         nil,
			"comments":     []int{17},
		},
		{
			"id":           17,
			"user_id":      8,
			"parent_id":    16,
			"community_id": 1,
			"content":      "Absolutely! Accessibility and collaboration among academia and the private sector will drive innovation. The future of technology is exciting!",
			"tags":         nil,
			"comments":     nil,
		},
		{
			"id":           18,
			"user_id":      11,
			"parent_id":    nil,
			"community_id": 2,
			"title":        "I just watched a documentary on autonomous robotic surgery.",
			"content":      "The precision and efficiency were mind-blowing! This is the future of medical robotics.",
			"tags":         []int{4, 5},
			"comments":     []int{19},
		},
		{
			"id":           19,
			"user_id":      12,
			"parent_id":    18,
			"community_id": 2,
			"content":      "It's incredible to think how far we've come. Do you know the name of the company behind this? I'd love to read more about their technology.",
			"tags":         nil,
			"comments":     []int{20},
		},
		{
			"id":           20,
			"user_id":      13,
			"parent_id":    19,
			"community_id": 2,
			"content":      "Yes, the company is called MedRobotics. They're making waves in the field of autonomous surgery. I attended a webinar where they discussed their latest research.",
			"tags":         nil,
			"comments":     []int{21},
		},
		{
			"id":           21,
			"user_id":      11,
			"parent_id":    20,
			"community_id": 2,
			"content":      "Thanks for the info! It's amazing to see real-world applications of AI in critical areas like healthcare. This could revolutionize the way surgeries are performed.",
			"tags":         nil,
			"comments":     []int{22},
		},
		{
			"id":           22,
			"user_id":      14,
			"parent_id":    21,
			"community_id": 2,
			"content":      "Definitely! Imagine remote surgeries where top surgeons can assist or even lead procedures from anywhere in the world. The potential for saving lives is enormous!",
			"tags":         nil,
			"comments":     nil,
		},
		{
			"id":           23,
			"user_id":      15,
			"parent_id":    nil,
			"community_id": 2,
			"content":      "Has anyone here experimented with open-source robotics platforms like ROS (Robot Operating System)? I'm starting a new project and looking for insights. #RoboRevolution #OpenSourceRobotics",
			"tags":         []int{4},
			"comments":     []int{24},
		},
		{
			"id":           24,
			"user_id":      16,
			"parent_id":    23,
			"community_id": 2,
			"content":      "I've worked with ROS a fair amount, mostly in an educational setting. It's a powerful tool, but it can be a bit complex. What kind of project are you working on?",
			"tags":         nil,
			"comments":     []int{25},
		},
		{
			"id":           25,
			"user_id":      17,
			"parent_id":    24,
			"community_id": 2,
			"content":      "ROS is excellent for developing robotic applications. It's got a bit of a learning curve, but the community support is robust. Feel free to ask any specific questions!",
			"tags":         nil,
			"comments":     []int{26},
		},
		{
			"id":           26,
			"user_id":      15,
			"parent_id":    25,
			"community_id": 2,
			"content":      "I'm developing a robot that can assist the elderly at home. Looking to integrate voice commands and some level of AI learning to understand the user's needs better.",
			"tags":         nil,
			"comments":     []int{27},
		},
		{
			"id":           27,
			"user_id":      18,
			"parent_id":    26,
			"community_id": 2,
			"content":      "That's a fantastic initiative! I've seen some work done with natural language processing and robotics. Maybe looking into existing AI voice platforms could speed up your development?",
			"tags":         nil,
			"comments":     nil,
		},
		{
			"id":           28,
			"user_id":      10,
			"parent_id":    nil,
			"community_id": 3,
			"title":        "I've recently switched to using reusable bags and bottles.",
			"content":      "It feels so good to minimize plastic waste! Any other simple changes I can make for sustainable living?",
			"tags":         []int{4, 5},
			"comments":     []int{29},
		},
		{
			"id":        29,
			"user_id":   14,
			"parent_id": 28,
			"content":   "That's a great start! You might consider using energy-efficient appliances and LED bulbs. Also, composting kitchen waste can be a rewarding way to recycle nutrients back into the soil.",
			"tags":      nil,
			"comments":  []int{30},
		},
		{
			"id":        30,
			"user_id":   7,
			"parent_id": 29,
			"content":   "Don't forget about conserving water by fixing leaks and using water-efficient fixtures. Every small change counts in our journey towards eco-friendly living!",
			"tags":      nil,
			"comments":  []int{31},
		},
		{
			"id":        31,
			"user_id":   18,
			"parent_id": 30,
			"content":   "I love these ideas! Also, think about supporting sustainable brands and products. They might be slightly more expensive, but the quality and ethical approach make it worth it.",
			"tags":      nil,
			"comments":  []int{32},
		},
		{
			"id":        32,
			"user_id":   21,
			"parent_id": 31,
			"content":   "True! It's an investment in both ourselves and our planet. There's so much we can do to live sustainably without sacrificing convenience or comfort.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":           33,
			"user_id":      23,
			"parent_id":    nil,
			"community_id": 3,
			"title":        "Has anyone here tried setting up solar panels at home?",
			"content":      "I'm interested in renewable energy but unsure where to start. Also, any tips for urban gardening?",
			"tags":         []int{6, 7},
			"comments":     []int{34},
		},
		{
			"id":        34,
			"user_id":   32,
			"parent_id": 33,
			"content":   "I installed solar panels last year, and it's been fantastic! Make sure to consult with a professional to understand your energy needs. As for urban gardening, vertical gardens and container planting work well in limited spaces.",
			"tags":      nil,
			"comments":  []int{35},
		},
		{
			"id":        35,
			"user_id":   40,
			"parent_id": 34,
			"content":   "Solar energy is a great investment. Don't forget about government incentives that might be available in your area. Urban gardening is an exciting hobby; you can grow herbs and vegetables even on a small balcony.",
			"tags":      nil,
			"comments":  []int{36},
		},
		{
			"id":        36,
			"user_id":   45,
			"parent_id": 35,
			"content":   "The sense of accomplishment from producing your energy and growing your food is priceless. Start small, do your research, and enjoy the process of learning and creating something sustainable.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":           37,
			"user_id":      6,
			"parent_id":    nil,
			"community_id": 4,
			"title":        "I've started an urban garden on my rooftop, using eco-friendly practices like composting and drip irrigation",
			"content":      "I'm amazed at how much food we can grow in a small space! Any other urban gardeners here?",
			"tags":         []int{7, 4},
			"comments":     []int{38},
		},
		{
			"id":        38,
			"user_id":   7,
			"parent_id": 37,
			"content":   "That's fantastic! I've been urban gardening for a couple of years now, and it's truly rewarding. The use of eco-friendly techniques makes it even better. Do you have any favorite plants or veggies you've grown?",
			"tags":      nil,
			"comments":  []int{39},
		},
		{
			"id":        39,
			"user_id":   8,
			"parent_id": 38,
			"content":   "I'm a beginner in urban gardening, but I've successfully grown herbs like basil and mint. It's lovely to have fresh herbs at hand, and knowing that it's done in an eco-friendly way makes it more special.",
			"tags":      nil,
			"comments":  []int{40},
		},
		{
			"id":        40,
			"user_id":   9,
			"parent_id": 39,
			"content":   "Herbs are a great start! Don't hesitate to experiment with other plants too. Urban gardening is not only about growing food but also building community and sustainability. Keep up the good work!",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":           41,
			"user_id":      10,
			"parent_id":    nil,
			"community_id": 4,
			"title":        "I've been using recycled containers for my potted plants and incorporating other recycled materials in my garden design.",
			"content":      "It's not only sustainable but adds a creative touch! How do you folks incorporate recycling in your urban gardens?",
			"tags":         []int{5, 6},
			"comments":     []int{42},
		},
		{
			"id":        42,
			"user_id":   11,
			"parent_id": 41,
			"content":   "I love that idea! I've used old wooden pallets to create vertical gardens. It's amazing how these recycled materials can be turned into something beautiful and functional.",
			"tags":      nil,
			"comments":  []int{43},
		},
		{
			"id":        43,
			"user_id":   12,
			"parent_id": 42,
			"content":   "Recycling in the garden is a brilliant way to be sustainable. I've used old tires for planting, and they look great. Also, composting kitchen waste is another way to recycle and enrich the soil.",
			"tags":      nil,
			"comments":  []int{44},
		},
		{
			"id":        44,
			"user_id":   13,
			"parent_id": 43,
			"content":   "I agree, composting is a game-changer! It's great for the environment and for our gardens. Urban gardening is indeed a creative and sustainable way to live. Let's keep sharing ideas and growing together!",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        45,
			"user_id":   8,
			"parent_id": nil,
			"title":     "Just baked my first sourdough bread!",
			"content":   "The texture and flavor are incredible. Anyone has tips for maintaining a healthy sourdough starter?",
			"tags":      []int{9, 10},
			"comments":  []int{46},
		},
		{
			"id":        46,
			"user_id":   9,
			"parent_id": 57,
			"content":   "I feed mine every day with equal parts water and flour. It's been thriving for months. Happy baking!",
			"tags":      nil,
			"comments":  []int{49},
		},
		{
			"id":        47,
			"user_id":   10,
			"parent_id": 46,
			"content":   "Do you use whole wheat or all-purpose flour? I've had success with rye flour too.",
			"tags":      nil,
			"comments":  []int{50},
		},
		{
			"id":        48,
			"user_id":   8,
			"parent_id": 47,
			"content":   "I used whole wheat flour. It adds more flavor. Thanks for the tips, I'll try rye flour next time.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        49,
			"user_id":   10,
			"parent_id": nil,
			"title":     "I'm exploring vegan recipes from around the world.",
			"content":   "Just tried a vegan Thai curry, and it was delicious! Any recommendations for other global vegan dishes?",
			"tags":      []int{11, 13},
			"comments":  []int{50, 51},
		},
		{
			"id":        50,
			"user_id":   11,
			"parent_id": 49,
			"content":   "You should try Ethiopian lentil stew. It's rich, flavorful, and naturally vegan!",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        51,
			"user_id":   14,
			"parent_id": 49,
			"content":   "Indian cuisine has many vegan options too. Have you tried Chana Masala?",
			"tags":      nil,
			"comments":  []int{52},
		},
		{
			"id":        52,
			"user_id":   10,
			"parent_id": 51,
			"content":   "Chana Masala sounds great! I'll give it a try. Thanks for the suggestion.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        53,
			"user_id":   12,
			"parent_id": nil,
			"title":     "Air frying is becoming a trend in home cooking.",
			"content":   "I've tried it with sweet potato fries and chicken wings. What other dishes are great for air frying?",
			"tags":      []int{12, 10},
			"comments":  []int{54, 55},
		},
		{
			"id":        54,
			"user_id":   13,
			"parent_id": 53,
			"content":   "I love air frying Brussels sprouts and salmon. It's a quick and healthy way to cook!",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        55,
			"user_id":   15,
			"parent_id": 53,
			"content":   "Air frying doughnuts is also amazing. They come out crispy and less greasy.",
			"tags":      nil,
			"comments":  []int{56},
		},
		{
			"id":        56,
			"user_id":   12,
			"parent_id": 55,
			"content":   "Doughnuts in the air fryer? I never thought of that. Definitely trying it this weekend!",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        57,
			"user_id":   16,
			"parent_id": nil,
			"title":     "Exploring Japanese cuisine this week.",
			"content":   "The artistry in sushi preparation is mesmerizing. Anyone has experience with making sushi at home?",
			"tags":      []int{11, 12, 9},
			"comments":  []int{58, 59},
		},
		{
			"id":        58,
			"user_id":   17,
			"parent_id": 57,
			"content":   "I've tried making sushi at home. It's all about having the right tools and fresh ingredients. Practice makes perfect!",
			"tags":      nil,
			"comments":  []int{59, 60},
		},
		{
			"id":        59,
			"user_id":   16,
			"parent_id": 58,
			"content":   "Did you take any online courses or just follow recipes? I'm keen to learn the right techniques.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        60,
			"user_id":   17,
			"parent_id": 58,
			"content":   "I watched several YouTube tutorials and read a few blogs. It's quite rewarding once you get the hang of it.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        59,
			"user_id":   18,
			"parent_id": 57,
			"content":   "Invest in a good knife and bamboo mat. It makes a huge difference in the sushi-making process.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        60,
			"user_id":   19,
			"parent_id": nil,
			"title":     "Dined at a gourmet vegan restaurant last night.",
			"content":   "The creativity in plant-based fine dining is astounding. Anyone here has favorite vegan gourmet recipes?",
			"tags":      []int{13, 9, 14},
			"comments":  []int{61, 62},
		},
		{
			"id":        61,
			"user_id":   20,
			"parent_id": 60,
			"content":   "I love making vegan risotto with wild mushrooms. It's rich and creamy without any dairy!",
			"tags":      nil,
			"comments":  []int{62, 63},
		},
		{
			"id":        62,
			"user_id":   19,
			"parent_id": 61,
			"content":   "That sounds delightful! Could you share the recipe?",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        63,
			"user_id":   20,
			"parent_id": 61,
			"content":   "Sure! I'll send you a link to the recipe. It's quite simple but tastes luxurious.",
			"tags":      nil,
			"comments":  []int{64},
		},
		{
			"id":        64,
			"user_id":   19,
			"parent_id": 63,
			"content":   "Thank you! I'll give it a try this weekend.",
			"comments":  nil,
			"tags":      nil,
		},
		{
			"id":        62,
			"user_id":   21,
			"parent_id": 60,
			"content":   "Vegan gourmet cooking is indeed a fascinating subject. I recently made a vegan chocolate tart that was a hit at a dinner party.",
			"comments":  nil,
			"tags":      nil,
		},
		/**/
		{
			"id":        63,
			"user_id":   10,
			"parent_id": nil,
			"title":     "Just finished reading '1984' by George Orwell",
			"content":   "The dystopian theme is as relevant today as it was when it was written. Thoughts?",
			"tags":      []int{20, 21},
			"comments":  []int{64},
		},
		{
			"id":        64,
			"user_id":   11,
			"parent_id": 63,
			"content":   "I agree, '1984' continues to be a cautionary tale. Have you read 'Brave New World' by Aldous Huxley? It offers another perspective on dystopian society.",
			"tags":      nil,
			"comments":  []int{65},
		},
		{
			"id":        65,
			"user_id":   12,
			"parent_id": 64,
			"content":   "Both of those works are seminal in understanding dystopian literature. They explore control, manipulation, and the human spirit.",
			"tags":      nil,
			"comments":  []int{66},
		},
		{
			"id":        66,
			"user_id":   10,
			"parent_id": 65,
			"content":   "I find the parallels between those fictional worlds and modern society quite unsettling. Literature can indeed be a mirror reflecting our own world.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        67,
			"user_id":   13,
			"parent_id": nil,
			"content":   "I'm looking to dive into fantasy novels. Any recommendations for a newbie?",
			"tags":      []int{22, 21},
			"comments":  []int{68, 70},
		},
		{
			"id":        68,
			"user_id":   14,
			"parent_id": 67,
			"content":   "You can't go wrong with 'The Hobbit' by J.R.R. Tolkien. It's a fantastic introduction to the genre.",
			"tags":      nil,
			"comments":  []int{69},
		},
		{
			"id":        69,
			"user_id":   15,
			"parent_id": 68,
			"content":   "Also, consider the 'Harry Potter' series by J.K. Rowling. It's engaging for readers of all ages.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        70,
			"user_id":   16,
			"parent_id": 67,
			"content":   "'The Chronicles of Narnia' by C.S. Lewis is also a wonderful choice. It offers a mix of adventure and profound themes.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        71,
			"user_id":   17,
			"parent_id": nil,
			"content":   "What's your favorite non-fiction book? I'm in the mood to learn something new!",
			"tags":      []int{23, 21},
			"comments":  []int{72},
		},
		{
			"id":        72,
			"user_id":   18,
			"parent_id": 71,
			"content":   "Sapiens: A Brief History of Humankind by Yuval Noah Harari is an insightful read. It gives an overview of the history of our species.",
			"tags":      nil,
			"comments":  []int{73},
		},
		{
			"id":        73,
			"user_id":   19,
			"parent_id": 72,
			"content":   "If you're interested in science, 'A Brief History of Time' by Stephen Hawking is an intriguing look at the universe.",
			"tags":      nil,
			"comments":  []int{74},
		},
		{
			"id":        74,
			"user_id":   17,
			"parent_id": 73,
			"content":   "Thanks for the recommendations! I'll check both of them out.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        75,
			"user_id":   20,
			"parent_id": nil,
			"content":   "Has anyone read 'Percy Jackson & The Olympians'? It's a great series for those interested in mythology.",
			"tags":      []int{24, 25},
			"comments":  []int{84},
		},
		{
			"id":        76,
			"user_id":   21,
			"parent_id": 75,
			"content":   "I love that series! It got me interested in Greek myths. Rick Riordan has a way of making these ancient stories come alive.",
			"tags":      nil,
			"comments":  []int{77},
		},
		{
			"id":        77,
			"user_id":   22,
			"parent_id": 76,
			"content":   "The 'Heroes of Olympus' series by the same author continues the adventure. It's a must-read if you liked 'Percy Jackson'.",
			"tags":      nil,
			"comments":  []int{78},
		},
		{
			"id":        78,
			"user_id":   23,
			"parent_id": 77,
			"content":   "I read them all! They are perfect for young readers. It made me curious about other mythologies too.",
			"tags":      nil,
			"comments":  nil,
		},
		{
			"id":        79,
			"user_id":   24,
			"parent_id": nil,
			"content":   "Looking for some age-appropriate mystery novels for my little sister. Any recommendations?",
			"tags":      []int{23, 25},
			"comments":  []int{88},
		},
		{
			"id":        80,
			"user_id":   25,
			"parent_id": 79,
			"content":   "'The Secret Seven' series by Enid Blyton is a classic choice. It's suitable for young readers and filled with intriguing mysteries.",
			"tags":      nil,
			"comments":  []int{89},
		},
		{
			"id":        81,
			"user_id":   26,
			"parent_id": 80,
			"content":   "Don't forget 'Nancy Drew' and 'The Hardy Boys'. They are timeless and provide engaging puzzles for young minds.",
			"tags":      nil,
			"comments":  []int{82},
		},
		{
			"id":        82,
			"user_id":   24,
			"parent_id": 81,
			"content":   "Thank you all! These will keep her entertained for sure. I might even read some myself!",
			"tags":      nil,
			"comments":  nil,
		},
	}
	communityList = []map[string]any{
		{
			"title":       "Innovators Hub",
			"slug":        "innovators-hub",
			"description": "A space for tech enthusiasts, engineers, and gadget lovers to come together and explore the latest innovations in the tech world. From cutting-edge technologies to in-depth gadget reviews, join the discussion and stay ahead of the curve.",
			"topics":      []int{1},
			"comments":    []int{1, 7, 13},
		},
		{
			"title":       "Future of Robotics",
			"slug":        "future-of-robotics",
			"description": "A dedicated group for those fascinated by robotics and artificial intelligence. Whether you're a researcher, hobbyist, or just curious about the field, this community brings together minds interested in the development and ethical implications of robots in our daily lives.",
			"topics":      []int{1},
			"comments":    []int{18, 23},
		},
		{
			"title":       "Eco Warriors",
			"slug":        "eco-warriors",
			"description": "A community dedicated to those passionate about the environment and sustainable living. Share eco-friendly tips, renewable energy innovations, and connect with like-minded individuals striving to make a greener world.",
			"topics":      []int{2},
			"comments":    []int{28, 33, 37},
		},
		{
			"title":       "Urban Gardening Enthusiasts",
			"slug":        "urban-gardening-enthusiasts",
			"description": "For lovers of urban gardening, this group fosters discussion on cultivating plants in city spaces. Share ideas on container gardening, rooftop gardens, balcony greenery, and more. Grow your green thumb in the heart of the city!",
			"topics":      []int{2},
			"comments":    []int{41, 45},
		},
		{
			"title":       "Home Cooks Unite",
			"slug":        "home-cooks-unite",
			"description": "A community for home cooking enthusiasts. Share your favorite recipes, cooking tips, and kitchen hacks. Whether you are a seasoned chef or just starting, join us in the joy of cooking at home.",
			"topics":      []int{3},
			"comments":    []int{57, 61, 65},
		},
		{
			"title":       "Global Gourmet Explorers",
			"slug":        "global-gourmet-explorers",
			"description": "Explore the flavors of the world. Discuss different cuisines, traditional dishes, and culinary techniques from across the globe. Travel through taste and share your gastronomic adventures here.",
			"topics":      []int{3},
			"comments":    []int{69, 74},
		},
		{
			"title":       "Mind and Body Balance",
			"slug":        "mind-and-body-balance",
			"description": "A community focused on integrating mental and physical well-being through mindfulness practices, holistic healing methods, and more.",
			"topics":      []int{4},
			"comments":    []int{75, 79, 83},
		},
		{
			"title":       "Active Life Advocates",
			"slug":        "active-life-advocates",
			"description": "For those who seek an active and healthy lifestyle through regular fitness routines, nutritional eating, and a positive mindset.",
			"topics":      []int{4},
			"comments":    []int{86, 90},
		},

		{
			"title":       "Literary Explorers",
			"slug":        "literary-explorers",
			"description": "A community for those who love to explore different genres and eras of literature. From historical novels to futuristic sci-fi, this group dives into all aspects of the literary world.",
			"topics":      []int{5},
			"comments":    []int{71, 75, 79},
		},
		{
			"title":       "Young Readers Club",
			"slug":        "young-readers-club",
			"description": "A dedicated group for young adults and those who love reading young adult fiction. Share your favorite coming-of-age stories, fantasy adventures, and romance novels.",
			"topics":      []int{5},
			"comments":    []int{83, 87},
		},
	}
	tagList = []map[string]string{
		{"name": "TechInnovation"},
		{"name": "GadgetGuru"},
		{"name": "FutureTech"},
		{"name": "RoboRevolution"},
		{"name": "AIInsights"},

		{"name": "SustainableLiving"},
		{"name": "EcoFriendly"},
		{"name": "RenewableEnergy"},
		{"name": "UrbanGardening"},
		{"name": "Recycling"},

		{"name": "CulinaryArt"},
		{"name": "HomeCooking"},
		{"name": "WorldCuisine"},
		{"name": "FoodTrends"},
		{"name": "VeganChoices"},

		{"name": "Mindfulness"},
		{"name": "HealthyEating"},
		{"name": "FitnessRoutine"},
		{"name": "HolisticHealing"},
		{"name": "MentalHealthAwareness"},

		{"name": "ClassicLiterature"},
		{"name": "ContemporaryFiction"},
		{"name": "FantasyLovers"},
		{"name": "NonFiction"},
		{"name": "YoungAdult"},
	}
	topicList = []map[string]string{
		{"name": "Tech Talk", "desc": "Discussions on the latest in technology and gadgets."},
		{"name": "Green Living", "desc": "Sustainable practices, renewable energy, and eco-friendly tips."},
		{"name": "Foodies' Nest", "desc": "All things culinary, from recipes to restaurant reviews."},
		{"name": "Wellness Wing", "desc": "Health and wellness tips, fitness, mental health discussions."},
		{"name": "Book Birds", "desc": "Book reviews, reading recommendations, and literary discussions."},
		{"name": "Sports in the Sky", "desc": "Posts about various sports, athletes, games, and tournaments."},
		{"name": "Business Beats", "desc": "Discussions on business trends, startups, and entrepreneurship."},
		{"name": "Movie Perch", "desc": "Movie reviews, trailers, celebrity news, and more."},
		{"name": "Travel Trails", "desc": "Travel experiences, tips, destination features, and cultural insights."},
		{"name": "Music Melodies", "desc": "Music reviews, artist features, and discussions on various genres."},
		{"name": "Artistic Altitudes", "desc": "Showcasing art, photography, digital design, etc."},
		{"name": "Science and Space", "desc": "Discoveries, news, and discussions about science and space."},
		{"name": "Gaming Grove", "desc": "Video game reviews, gaming news, eSports, and game development."},
		{"name": "Fashion Flight", "desc": "Latest fashion trends, style tips, designer features."},
		{"name": "Pet Park", "desc": "Posts about pet care, cute pet photos, animal welfare discussions."},
		{"name": "DIY Den", "desc": "Home improvement tips, crafting, DIY project ideas."},
		{"name": "Education Elevations", "desc": "Topics around education, online learning, tutorials."},
		{"name": "Cultural Crosswinds", "desc": "Posts about different cultures, traditions, and societal issues."},
		{"name": "Home & Hearth", "desc": "Topics related to home decor, family, parenting, and gardening."},
		{"name": "Financial Flight", "desc": "Discussions about personal finance, investing, and economics."},
	}
}
