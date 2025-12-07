package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/vanhieuhp/social/internal/store"
)

var usernames = []string{
	"OceanNova", "TechHorizon", "PixelRush", "CodeVortex", "SilentComet",
	"NeonRider", "ByteHunter", "LunarEcho", "AquaGlide", "CyberSparrow",

	"StormPulse", "RogueMatrix", "ZenCrafter", "ShadowBinary", "WaveSynth",
	"FrostCircuit", "SolarMender", "CrimsonLogic", "IronPetal", "NightVoyager",

	"QuantumShift", "BlueInferno", "SwiftRaptor", "EchoBlade", "MidnightNova",
	"CrystalGear", "StaticRunner", "BlazeDrifter", "BinaryNomad", "VoidWalker",

	"HyperGlyph", "StarFlicker", "GhostCoder", "NebulaDash", "ThunderMint",
	"RunePilot", "ChromePulse", "EchoRanger", "SkyForge", "PulseSeeker",

	"ArcticShade", "NovaSprinter", "SilentRidge", "CircuitGlide", "NightSpark",
	"ByteWanderer", "LunarBreaker", "StormCrafter", "SteelQuasar", "CloudStinger",
}

var titles = []string{
	"Boost Your Productivity Today",
	"Why Simple Habits Matter",
	"Mastering Focus in a Digital World",
	"Top Tools for Daily Efficiency",
	"How to Build Strong Routines",
	"The Power of Small Improvements",
	"Reduce Stress with Easy Steps",
	"Quick Tips for Better Time Management",
	"How to Stay Motivated Long-Term",
	"Lessons from Successful People",
	"Healthy Morning Rituals",
	"How to Avoid Burnout",
	"Small Wins That Change Everything",
	"Making Better Daily Decisions",
	"The Art of Staying Consistent",
	"Improve Your Life with Minimal Effort",
	"Why Planning Saves You Time",
	"Unlocking Your Creative Potential",
	"How to Build Momentum Fast",
	"Simple Hacks for a Better Week",
}

var contents = []string{
	"Discover simple habits that help you stay productive without feeling overwhelmed.",
	"Learn why tiny daily improvements can lead to long-term success and stability.",
	"Find practical strategies to maintain deep focus in a world full of distractions.",
	"Explore essential tools that make your workflow faster and more efficient.",
	"Build strong routines that support your goals and reduce decision fatigue.",
	"Understand how small changes can create a big impact on your personal growth.",
	"Reduce stress with easy, actionable steps you can apply immediately.",
	"Improve your time management with quick tips that fit into any schedule.",
	"Stay motivated by applying simple techniques used by high achievers.",
	"Discover life lessons from successful people and how to apply them daily.",
	"Start your day right with healthy morning rituals that boost your energy.",
	"Learn how to recognize and prevent burnout before it affects your work.",
	"Celebrate small wins to stay consistent and build long-term momentum.",
	"Make better choices every day using a simple decision-making framework.",
	"Stay consistent with a practical approach that helps you avoid procrastination.",
	"Improve your life with minimal-effort hacks that save time and stress.",
	"See how proper planning can help you achieve more in less time.",
	"Unlock your creativity through easy exercises and mindset shifts.",
	"Build momentum quickly by taking small but intentional actions.",
	"Use these simple weekly habits to create a happier and more balanced life.",
}

var tags = []string{
	"productivity",
	"motivation",
	"self-improvement",
	"focus",
	"habits",
	"lifestyle",
	"mindset",
	"time-management",
	"success",
	"workflow",
	"health",
	"routine",
	"creativity",
	"minimalism",
	"planning",
	"consistency",
	"personal-growth",
	"wellness",
	"stress-management",
	"daily-tips",
}

var commentSamples = []string{
	"Great post!",
	"Very helpful!",
	"Thanks for sharing!",
	"Nice tips!",
	"Love this!",
	"So true!",
	"Well explained!",
	"Very inspiring!",
	"Helpful info!",
	"Awesome content!",
	"Good point!",
	"Amazing!",
	"Simple and clear!",
	"I agree!",
	"Thanks for the insight!",
	"Nice reminder!",
	"So useful!",
	"Brilliant!",
	"Well said!",
	"Really helpful!",
	"Keep it up!",
	"Great advice!",
	"This is gold!",
	"Very motivating!",
	"Good to know!",
	"Super helpful!",
	"Nice one!",
	"Thanks a lot!",
	"Short and sweet!",
	"Love the simplicity!",
	"Very practical!",
	"Good read!",
	"Perfect!",
	"This helps!",
	"Interesting!",
	"I needed this!",
	"Well written!",
	"Great insight!",
	"Very clear!",
	"Totally agree!",
	"Solid advice!",
	"Spot on!",
	"Good reminder!",
	"Useful tips!",
	"Thanks for this!",
	"Great explanation!",
	"Nicely done!",
	"This is helpful!",
	"Love the message!",
}

func Seed(store store.Storage) {
	ctx := context.Background()
	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error create user: ", err)
			return
		}
	}

	posts := generatePost(100, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error create post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.CreateComment(ctx, comment); err != nil {
			log.Println("Error create comment: ", err)
			return
		}
	}

	log.Println("Seeding data complete!")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d@example.com", i),
			Password: "123456",
		}
	}

	return users
}

func generatePost(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]
		comments[i] = &store.Comment{
			Content: commentSamples[rand.Intn(len(commentSamples))],
			PostId:  post.ID,
			UserId:  user.ID,
		}
	}

	return comments
}
