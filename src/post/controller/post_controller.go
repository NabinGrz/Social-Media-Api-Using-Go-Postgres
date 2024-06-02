package postController

import (
	"net/http"

	userPostModel "github.com/NabinGrz/SocialMedia/src/post/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// func GetPostDetails(c *gin.Context, db *gorm.DB) {

// 	// cursor, err := dbConnection.PostCollection.Aggregate(context.Background(), getAllPostPipeline())
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }
// 	// defer cursor.Close(context.Background())
// 	// var allPost []userPostModel.SocialMediaPost
// 	// for cursor.Next(c) {
// 	// 	var post postModel.SocialMediaPost

// 	// 	if err := cursor.Decode(&post); err != nil {
// 	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 		return
// 	// 	}
// 	// 	allPost = append(allPost, post)
// 	// }

// 	// if err := cursor.Err(); err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }

// 	// c.JSON(http.StatusOK, allPost)
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	var post userPostModel.SocialMediaPost
// 	// p, _ := uuid.Parse("7b427d94-5ee0-44e8-9b1d-e189682823be")
// 	//.Preload("MediaDetails").Preload("CommentDetails")
// 	if err := db.Preload("User").First(&post, "id = ?", id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, post)

// }

func GetPostDetails(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var post userPostModel.SocialMediaPost
	if err := db.Preload("User").First(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context, db *gorm.DB) {

	// id, _ := strconv.Atoi(c.Param("id"))
	var post userPostModel.SocialMediaPost
	userID := c.GetString("userid")
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedUUID, _ := uuid.Parse(userID)
	post.UserID = parsedUUID

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}

// func UpdatePost(c *gin.Context) {
// 	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))

// 	var updatedPost postModel.SocialMediaPost
// 	var foundPost postModel.SocialMediaPost

// 	id := c.Param("id")

// 	objID, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objID}

// 	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

// 	if result.Err() != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}

// 	err := c.ShouldBindJSON(&updatedPost)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	result.Decode(&foundPost)

// 	if userID != foundPost.User {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "You cannot update others post"})
// 		return
// 	}
// 	update := bson.M{"$set": bson.M{"caption": updatedPost.Caption}}
// 	updateResult, _ := dbConnection.PostCollection.UpdateMany(context.Background(), filter, update)

// 	if updateResult.MatchedCount == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
// }
// func DeletePost(c *gin.Context) {
// 	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))

// 	var post postModel.SocialMediaPost

// 	id := c.Param("id")

// 	objID, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objID}

// 	err := dbConnection.PostCollection.FindOne(context.Background(), filter).Decode(&post)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	if userID != post.User {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "You cannot delete others post"})
// 		return
// 	}
// 	result, _ := dbConnection.PostCollection.DeleteOne(context.Background(), filter)

// 	if result.DeletedCount == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
// }

// func LikePost(c *gin.Context) {
// 	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
// 	var foundPost postModel.SocialMediaPost

// 	id := c.Param("id")

// 	objID, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objID}

// 	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

// 	if result.Err() != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}
// 	result.Decode(&foundPost)
// 	update := bson.M{"$addToSet": bson.M{"likeby": userID}}
// 	updateResult, _ := dbConnection.PostCollection.UpdateMany(context.Background(), filter, update)
// 	fmt.Println(updateResult)

// 	c.JSON(http.StatusOK, gin.H{"message": "Post has been liked successfully"})
// }
// func SharePost(c *gin.Context) {
// 	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
// 	var foundPost postModel.SocialMediaPost

// 	id := c.Param("id")

// 	objID, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objID}

// 	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

// 	if result.Err() != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}
// 	result.Decode(&foundPost)
// 	update := bson.M{"$addToSet": bson.M{"shares": userID}}
// 	updateResult, _ := dbConnection.PostCollection.UpdateMany(context.Background(), filter, update)
// 	fmt.Println(updateResult)
// 	c.JSON(http.StatusOK, gin.H{"message": "Post has been shared successfully"})
// }
// func CommentPost(c *gin.Context) {
// 	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
// 	var foundPost postModel.SocialMediaPost
// 	var comment postModel.CommentDetail

// 	id := c.Param("id")

// 	objID, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objID}

// 	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

// 	if result.Err() != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}
// 	result.Decode(&foundPost)
// 	err := c.ShouldBindJSON(&comment)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if comment.Comment == "" {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{
// 			"error": "Please enter your comment",
// 		})
// 		return
// 	}
// 	comment.User = userID
// 	comment.Date = time.Now()
// 	comment.ID = primitive.NewObjectID()

// 	update := bson.M{"$push": bson.M{
// 		"comments": comment,
// 	}}
// 	updateResult, _ := dbConnection.PostCollection.UpdateOne(context.Background(), filter, update)
// 	fmt.Println(updateResult)
// 	c.JSON(http.StatusOK, gin.H{"message": "Post has been commented successfully"})
// }

// func getAllPostPipeline() mongo.Pipeline {
// 	pipeline := mongo.Pipeline{
// 		{{ // $lookup stage
// 			Key: "$lookup",
// 			Value: bson.D{
// 				{Key: "from", Value: "user"},
// 				{Key: "localField", Value: "user"},
// 				{Key: "foreignField", Value: "_id"},
// 				{Key: "as", Value: "userdata"},
// 			},
// 		}},
// 		{{ // $lookup stage
// 			Key: "$lookup",
// 			Value: bson.D{
// 				{Key: "from", Value: "user"},
// 				{Key: "localField", Value: "likeby"},
// 				{Key: "foreignField", Value: "_id"},
// 				{Key: "as", Value: "likeby"},
// 			},
// 		}},
// 		{{ // $lookup stage
// 			Key: "$lookup",
// 			Value: bson.D{
// 				{Key: "from", Value: "user"},
// 				{Key: "localField", Value: "shares"},
// 				{Key: "foreignField", Value: "_id"},
// 				{Key: "as", Value: "shares"},
// 			},
// 		}},
// 		// {{ // $lookup stage
// 		// 	Key: "$lookup",
// 		// 	Value: bson.D{
// 		// 		{Key: "from", Value: "user"},
// 		// 		{Key: "localField", Value: "comments.user"},
// 		// 		{Key: "foreignField", Value: "_id"},
// 		// 		{Key: "as", Value: "commentUsers"},
// 		// 	},
// 		// }},
// 		bson.D{{Key: "$unwind", Value: bson.D{
// 			{Key: "path", Value: "$commentUsers"},
// 			{Key: "preserveNullAndEmptyArrays", Value: true},
// 		}}},
// 		bson.D{{Key: "$addFields", Value: bson.D{
// 			{Key: "comments", Value: bson.D{
// 				{Key: "$cond", Value: bson.A{
// 					// Condition to check if comments field is null
// 					bson.D{
// 						{Key: "$eq", Value: bson.A{"$comments", nil}},
// 					},
// 					// If comments is null, set it to a default comment structure
// 					bson.A{bson.D{
// 						{Key: "_id", Value: ""},
// 						{Key: "comment", Value: ""},
// 						{Key: "date", Value: ""},
// 						{Key: "commentUsers", Value: bson.A{}},
// 						{Key: "user", Value: ""},
// 					}},
// 					// If comments is not null, map each comment
// 					bson.D{
// 						{Key: "$map", Value: bson.D{
// 							{Key: "input", Value: "$comments"},
// 							{Key: "as", Value: "comment"},
// 							{Key: "in", Value: bson.D{
// 								{Key: "_id", Value: bson.D{
// 									{Key: "$convert", Value: bson.D{
// 										{Key: "input", Value: "$$comment._id"},
// 										{Key: "to", Value: "objectId"},
// 									}},
// 								}},
// 								{Key: "comment", Value: "$$comment.comment"},
// 								{Key: "date", Value: bson.D{{Key: "$toDate", Value: "$$comment.date"}}}, // Convert date to Date type
// 								{Key: "commentUsers", Value: "$$comment.commentUsers"},
// 								{Key: "user", Value: bson.D{{Key: "$toObjectId", Value: "$$comment.user"}}},
// 							}},
// 						}},
// 					},
// 				}},
// 			}},
// 		}}},

// 		bson.D{{Key: "$project", Value: bson.D{
// 			{Key: "_id", Value: 1},
// 			{Key: "caption", Value: 1},
// 			{Key: "user", Value: 1},
// 			{Key: "date", Value: bson.D{{Key: "$toDate", Value: "$date"}}},
// 			{Key: "media", Value: 1},
// 			{Key: "likeby", Value: 1},
// 			{Key: "shares", Value: 1},
// 			{Key: "comments", Value: 1},
// 			{Key: "userdata", Value: 1},
// 		}}},
// 	}
// 	return pipeline
// }
