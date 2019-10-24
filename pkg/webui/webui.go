package webui

import (
	"errors"
	"net/http"

	k8scli "github.com/akishitara/cronjob-operator/pkg/client"

	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		cronJobs := k8scli.ActiveCronjobsAllGet()
		c.HTML(http.StatusOK, "top.tmpl", gin.H{
			"title":    "CronJobChart",
			"cronJobs": cronJobs,
		})
	})
	router.GET("/timeline.js", func(c *gin.Context) {
		c.HTML(http.StatusOK, "timeline.js.tmpl", gin.H{
			"data": k8scli.ActiveCronjobsAllGet(),
		})
	})
	router.GET("/api/cronjobs/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, k8scli.ActiveCronjobsAllGet())
	})
	router.GET("/api/cronjobs/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, k8scli.ActiveCronjobsAllGet())
	})
	router.GET("/api/logs/:podName/json", func(c *gin.Context) {
		name := string(c.Param("podName"))
		c.JSON(http.StatusOK, k8scli.PodLogs(name))
	})
	router.GET("/logs/:podName", func(c *gin.Context) {
		cronJobs := k8scli.ActiveCronjobsAllGet()
		name := string(c.Param("podName"))
		c.HTML(http.StatusOK, "logs.tmpl", gin.H{
			"title":    "CronJobChart(Log)",
			"cronJobs": cronJobs,
			"log":      k8scli.PodLogs(name),
			"podName":  name,
		})
	})
	router.POST("/exec/:cronjobName", func(c *gin.Context) {
		name := string(c.Param("cronjobName"))
		k8scli.TriggerCronJob(name)
	})
	router.Run(":80")
	return errors.New("fail webui")
}
