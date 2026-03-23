package service

import (
	"fmt"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/email"
	"ms-tagpeak/pkg/logster"
)

func SendBugReport(report models.BugReportRequest) error {
	logster.StartFuncLog()

	recipient := dotenv.GetEnv("BUG_REPORT_RECIPIENT")

	prefix := "BUG REPORT"
	if report.Type != "technical" {
		prefix = "QUESTION"
	}
	subject := fmt.Sprintf("%s - %s (%s)", prefix, report.Name, report.Email)

	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>%s</h2>
			<table border="1" cellpadding="8" cellspacing="0" style="border-collapse: collapse;">
				<tr><td><strong>Name</strong></td><td>%s</td></tr>
				<tr><td><strong>Email</strong></td><td>%s</td></tr>
				<tr><td><strong>Type</strong></td><td>%s</td></tr>
				<tr><td><strong>Description</strong></td><td>%s</td></tr>
			</table>
		</body>
		</html>`,
		subject, report.Name, report.Email, report.Type, report.Description,
	)

	dto := email.SendRawEmailDTO{
		To:      recipient,
		Subject: subject,
		Body:    body,
		ReplyTo: report.Email,
	}

	if report.Attachment != nil {
		dto.Attachments = []email.SendEmailAttachment{
			{
				Filename: report.Attachment.Filename,
				Data:     report.Attachment.Data,
				MimeType: report.Attachment.MimeType,
			},
		}
	}

	_, errMap := email.SendRawEmail(dto)
	if errMap != nil {
		logster.Error(fmt.Errorf("%v", errMap), "Error sending bug report email")
		logster.EndFuncLog()
		return fmt.Errorf("failed to send bug report email")
	}

	logster.EndFuncLog()
	return nil
}
