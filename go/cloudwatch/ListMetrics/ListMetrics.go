/*
   Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
    "fmt"
    "strconv"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudwatch"
)

// GetMetrics gets the name, namespace, and dimension name of your Amazon CloudWatch metrics
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
// Output:
//     If successful, the metrics and nil
//     Otherwise, nil and an error from a call to ListMetrics
func GetMetrics(sess *session.Session) (*cloudwatch.ListMetricsOutput, error) {
    // Create CloudWatch client
    svc := cloudwatch.New(sess)

    result, err := svc.ListMetrics(nil)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func main() {
    // Initialize a session that the SDK uses to load
    // credentials from the shared credentials file ~/.aws/credentials
    // and configuration from the shared configuration file ~/.aws/config.
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    result, err := GetMetrics(sess)
    if err != nil {
        fmt.Println("Could not get metrics")
        return
    }

    fmt.Println("Metrics:")
    numMetrics := 0

    for _, m := range result.Metrics {
        fmt.Println("   Metric Name: " + *m.MetricName)
        fmt.Println("   Namespace:   " + *m.Namespace)
        fmt.Println("   Dimensions:")
        for _, d := range m.Dimensions {
            fmt.Println("      " + *d.Name + ": " + *d.Value)
        }

        fmt.Println("")
        numMetrics++
    }

    fmt.Println("Found " + strconv.Itoa(numMetrics) + " metrics")
}
