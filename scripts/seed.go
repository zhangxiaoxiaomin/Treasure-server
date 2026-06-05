package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"treasure-server/database"
	"treasure-server/models"
	"treasure-server/repository"
)

func main() {
	// Initialize database
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/treasure.db"
	}

	if err := database.Init(dbPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repo := repository.NewCollectionRepo()

	// Seed data matching the mini program's collections
	seedData := []models.CreateCollectionRequest{
		{
			ID:            "Qing-Porcelain",
			TitleCN:       "清代青花瓷瓶",
			TitleEN:       "Qing Dynasty Blue and White Porcelain",
			Category:      "porcelain",
			Image:         "https://lh3.googleusercontent.com/aida-public/AB6AXuDZdcK1OwA23C554FUo8tQeHPMMm0dUWEuPJO65ftPVL--otKd8qScL_pOGO5V77SBmjJlAl5Ft9CRd29deuvL7HKmQ94zCmKFW_UVwiD-EoOwQMtEZV1kEK2lxyYnDy_2j5ZO6FT0-WrvFX6fIdp3OWR1A-9Nsl0jlDKSb_rdfg8V6hE4ShSellMpy-2ROTEdX14_UtlCTrn2ABtdnuXjvQQ9zLT74zmqj-9bQPS9v7cF-e8hc8wVwbTyR6TE8KSUodYqRRkYOc78",
			DetailImages:  []string{"https://lh3.googleusercontent.com/aida-public/AB6AXuDZdcK1OwA23C554FUo8tQeHPMMm0dUWEuPJO65ftPVL--otKd8qScL_pOGO5V77SBmjJlAl5Ft9CRd29deuvL7HKmQ94zCmKFW_UVwiD-EoOwQMtEZV1kEK2lxyYnDy_2j5ZO6FT0-WrvFX6fIdp3OWR1A-9Nsl0jlDKSb_rdfg8V6hE4ShSellMpy-2ROTEdX14_UtlCTrn2ABtdnuXjvQQ9zLT74zmqj-9bQPS9v7cF-e8hc8wVwbTyR6TE8KSUodYqRRkYOc78"},
			Views:         1240,
			Likes:         340,
			CommentCount:  16,
			BadgeCN:       "实时竞拍",
			BadgeEN:       "Live Auction",
			DateStrCN:     "发布于 2024-05-20",
			DateStrEN:     "Posted on 2024-05-20",
			DescriptionCN: "胎质坚贞，釉色清朗莹润。青花发色深沉，以赶珠龙纹描绘，龙鳞细密，张牙舞爪，极富动感与神威。",
			DescriptionEN: "An exquisite close-up of a Qing dynasty porcelain vase with intricate blue and white patterns. The lighting is soft and focused, emphasizing historical craftsmanship.",
			DetailDescCN:  "此青花缠枝龙纹赏瓶为清代官窑杰作。器型庄重典雅，线条优美。画工细腻，青花发色浓翠艳丽，深入胎骨。所绘祥龙呼之欲出，体现了当时最高超的白瓷与青花烧制工艺。保存极其完好，具极高收藏附加值。",
			DetailDescEN:  "This blue and white dragon vase is a masterpiece of the imperial kiln of the Qing Dynasty. The shape is dignified, elegant, and perfectly proportioned. Painted with intense cobalt, the decoration depicts a dynamic five-clawed imperial dragon chasing the flaming pearl.",
			Comments: []models.Comment{
				{ID: "1", User: "高山流水", Avatar: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&h=100&fit=crop", Role: "资深鉴赏家", Time: "1天前", Content: "青花发色极其纯正，明显是典型的苏麻离青或精炼浙料，不可多得。"},
				{ID: "2", User: "收藏学徒", Avatar: "https://images.unsplash.com/photo-1544005313-94ddf0286df2?w=100&h=100&fit=crop", Time: "12小时前", Content: "这个制式真是太典雅了，细节和神韵无一不精。"},
			},
		},
		{
			ID:            "Ming-Scroll",
			TitleCN:       "皇家御制书法卷轴 / 明代书法卷轴",
			TitleEN:       "Imperial Era Calligraphy Scroll / Ming Dynasty Scroll",
			Category:      "calligraphy",
			Image:         "https://lh3.googleusercontent.com/aida-public/AB6AXuAk03U_5L0ofnKftNzrWL3s2ID40zFKhD7sGf4gk-LdTC-6hXIqjuhGUaq7qtqYtaQlkp8No9UIrOvpvmt9d8PMILYlsT29W1vkarq5Hn235zymDKzjw2cUfjdsLKcBqqH2tJ2dMLV2AlA3jDtdhlKUGeok-2jJBdZqcBJNtGhZUT2xKpABrmAX5NHIOkzMYk8biS929d97yDGMJEUwouO6p8Gtm-412ZAfg6eaYYMb-HmwvkvjGMqeVnODzUATfJAM8aMvCN-IbwU",
			DetailImages:  []string{"https://lh3.googleusercontent.com/aida-public/AB6AXuAk03U_5L0ofnKftNzrWL3s2ID40zFKhD7sGf4gk-LdTC-6hXIqjuhGUaq7qtqYtaQlkp8No9UIrOvpvmt9d8PMILYlsT29W1vkarq5Hn235zymDKzjw2cUfjdsLKcBqqH2tJ2dMLV2AlA3jDtdhlKUGeok-2jJBdZqcBJNtGhZUT2xKpABrmAX5NHIOkzMYk8biS929d97yDGMJEUwouO6p8Gtm-412ZAfg6eaYYMb-HmwvkvjGMqeVnODzUATfJAM8aMvCN-IbwU"},
			Views:         890,
			Likes:         156,
			CommentCount:  12,
			BadgeCN:       "即将开始",
			BadgeEN:       "Upcoming",
			DateStrCN:     "发布于 2024-05-18",
			DateStrEN:     "Posted on 2024-05-18",
			DescriptionCN: "书法精妙，带有历史印鉴。这件稀有的艺术品展示了16世纪晚期大师级的笔触，在私人收藏中保存完好。",
			DescriptionEN: "Exquisite calligraphy with historical stamps. This rare piece showcases the masterful brushwork of the late 16th century, preserved in remarkable condition.",
			DetailDescCN:  "此卷《明人书迹》为典型的十六世纪晚期行草书珍品。其纸质莹润，墨色入木三分，展现了明代文人书法的深厚底蕴与潇洒神采。作品可见起笔沉稳，收笔有力，转折处方圆兼备。落款与钤印位置考究，经多枚流传印鉴证实其在乾隆时期曾进入内府收藏，后流散于民间。纸张采用特制皮纸，虽历经五百年沧桑，依然韧性十足，且保留了当时特有的淡淡书香。",
			DetailDescEN:  "This handscroll of calligraphy captures the masterful brushmanship of the late Ming Dynasty. Done in running script, the ink strikes deep into the high-quality local mulberry paper. Features several collector seals indicating it was previously in the imperial collection during the Qianlong period. Truly an authoritative asset with rich historiographical value.",
			Comments: []models.Comment{
				{ID: "3", User: "翰墨斋主", Avatar: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100&h=100&fit=crop", Role: "书法协会顾问", Time: "2天前", Content: "笔力遒劲，字里行间透着一股魏晋风骨，这件珍藏品相完美！"},
			},
		},
		{
			ID:            "Song-Coins",
			TitleCN:       "宋代青铜古钱币",
			TitleEN:       "Song Dynasty Bronze Coinage",
			Category:      "coins",
			Image:         "https://images.unsplash.com/photo-1589758438368-0ad531db3366?w=800&auto=format&fit=crop&q=80",
			DetailImages:  []string{"https://images.unsplash.com/photo-1589758438368-0ad531db3366?w=800&auto=format&fit=crop&q=80"},
			Views:         2100,
			Likes:         412,
			CommentCount:  22,
			BadgeCN:       "正品鉴证",
			BadgeEN:       "Certified",
			DateStrCN:     "发布于 2024-05-15",
			DateStrEN:     "Posted on 2024-05-15",
			DescriptionCN: "北宋崇宁重宝，青铜质地。由于保存环境卓越，通体结满温润的绿锈，字口深峻，轮廓清晰。",
			DescriptionEN: "Chongning Heavily Treasure of Northern Song dynasty. Copper bronze casting. Coated with beautiful jade-like patina, crisp characters.",
			DetailDescCN:  "宋徽宗御书瘦金体崇宁重宝。字体健拔，铁画银钩，为中国古代钱币书法之巅峰。此枚古币郭肉均整，地章平整，传世包浆温润，无任何人工修补。是研究两宋铸币史及徽宗书法不可多得的实物史料。",
			DetailDescEN:  "This coin exhibits the legendary \"Slender Gold\" script written by Emperor Huizong himself. Known for its graceful, bone-thin yet intensely powerful strokes. This cast bronze specimen has survived with magnificent green-curry patina and sharp edge definition.",
			Comments: []models.Comment{
				{ID: "4", User: "泉界老李", Avatar: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=100&h=100&fit=crop", Role: "古泉名宿", Time: "3天前", Content: "瘦金体神韵十足，真是钱币艺术的最高巅峰。"},
			},
		},
		{
			ID:            "Han-Jade",
			TitleCN:       "汉代螭龙纹玉璧",
			TitleEN:       "Han Dynasty Chilong Jade Disc",
			Category:      "antiques",
			Image:         "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?w=800&auto=format&fit=crop&q=80",
			DetailImages:  []string{"https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?w=800&auto=format&fit=crop&q=80"},
			Views:         1450,
			Likes:         289,
			CommentCount:  9,
			BadgeCN:       "私人馆藏",
			BadgeEN:       "Private Collection",
			DateStrCN:     "发布于 2024-05-10",
			DateStrEN:     "Posted on 2024-05-10",
			DescriptionCN: "和田白玉材质，局部带褐色玉皮沁色。璧面精雕螭龙纹，刀工生动，矫健有力，极富张力。",
			DescriptionEN: "White Hetian nephrite jade disc featuring carved Chilong dragons. Exhibiting rich natural calcification and soil staining.",
			DetailDescCN:  "此汉代螭龙纹玉璧采用优质和田籽料雕琢，琢工遒劲有力，游丝描线细密流利。龙首高昂，体态盘曲，肌肉饱满。历经两千年洗礼，沁色入肌，玻璃光泽幽邃深沉，完美契合汉代治玉之雄浑大气的风格。",
			DetailDescEN:  "An important nephrite jade bi disc from the Han Dynasty. The masterfully low-relief Chilong coiled across the surface shows remarkable fluidity and three-dimensional definition. Highly lustrous, representing the finest jade carving work in antiquity.",
			Comments: []models.Comment{
				{ID: "5", User: "琢玉客", Avatar: "https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&h=100&fit=crop", Time: "5天前", Content: "玉质温润，沁色自然，线条极其流畅，毫无生涩感。"},
			},
		},
	}

	for _, data := range seedData {
		existing, _ := repo.GetByID(data.ID)
		if existing != nil {
			fmt.Printf("Collection %s already exists, skipping...\n", data.ID)
			continue
		}

		result, err := repo.Create(&data)
		if err != nil {
			log.Printf("Failed to create collection %s: %v", data.ID, err)
			continue
		}

		// Pretty print the created collection
		jsonBytes, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("Created collection: %s\n", string(jsonBytes))
	}

	fmt.Println("Seed completed!")
}