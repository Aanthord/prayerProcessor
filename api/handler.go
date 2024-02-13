// /prayerProcessor/api/handler.go

package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Aanthord/prayerProcessor/model"
    "github.com/Aanthord/prayerProcessor/service"
    "fmt"
)




// @Summary Process Prayer Request
// @Description Processes an incoming prayer request by performing PCA and sending the result to Kafka topics.
// @Tags prayer
// @Accept json
// @Produce json
// @Param prayer body model.Prayer true "Prayer Request"
// @Success 200 {object} map[string]interface{} "message: Prayer request processed successfully"
// @Failure 400 {object} map[string]interface{} "error: Could not parse prayer request"
// @Failure 500 {object} map[string]interface{} "error: Failed to process prayer request or send to Kafka"
// @Router /processPrayer [post]
func HandlePrayerRequest(kafkaProducer *service.KafkaProducer) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Parse the incoming JSON payload into the Prayer struct
        var prayer model.Prayer
        if err := c.BodyParser(&prayer); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Could not parse prayer request",
            })
        }

        // Process the prayer text through PCA
        pcaResult, err := service.ComputePCA(prayer.Text)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to process prayer request",
            })
        }

        // Convert PCA result to a string (simplified for example purposes)
        pcaResultStr := fmt.Sprintf("%v", pcaResult)

        // Send the processed data to the Kafka topics
        if err := kafkaProducer.SendMessage("processedForHeatMap", pcaResultStr); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to send to Kafka topic: processedForHeatMap",
            })
        }

        if err := kafkaProducer.SendMessage("recordToChain", pcaResultStr); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to send to Kafka topic: recordToChain",
            })
        }

        // Respond to the request indicating success
        return c.JSON(fiber.Map{
            "message": "Prayer request processed successfully",
        })
    }
}

