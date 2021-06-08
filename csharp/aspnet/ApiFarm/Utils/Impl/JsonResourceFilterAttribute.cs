using System;
using System.IO;
using System.Linq;
using System.Reflection;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// An attribute to check whether JSON request is valid and respond with custom message.
    /// </summary>
    public class JsonResourceFilterAttribute : Attribute, IResourceFilter
    {
        private static Func<PropertyInfo, string> getJsonName = propertyInfo =>
        {
            var jsonPropertyAttribute = propertyInfo.GetCustomAttributes(typeof(JsonPropertyAttribute), false).FirstOrDefault() as JsonPropertyAttribute;

            if (jsonPropertyAttribute is null)
            {
                return propertyInfo.Name.ToLower();
            }

            return jsonPropertyAttribute.PropertyName;
        };

        private Type targetModelType;

        /// <summary>
        /// Initializes a new instance of the <see cref="JsonResourceFilterAttribute"/> class.
        /// </summary>
        /// <param name="targetModelType">The model type to which the Json is being unmarshalled to.</param>
        public JsonResourceFilterAttribute(Type targetModelType)
        {
            this.targetModelType = targetModelType;
        }

        /// <summary>
        /// Called after the mvc pipeline.
        /// </summary>
        /// <param name="context">The pipeline process context.</param>
        public void OnResourceExecuted(ResourceExecutedContext context)
        {
        }

        /// <summary>
        /// Called after authorization filters. Validates JSON request body; responding with custom invalid JSON message.
        /// </summary>
        /// <param name="context">The pipeline process context.</param>
        public void OnResourceExecuting(ResourceExecutingContext context)
        {
            try
            {
                context.HttpContext.Request.EnableBuffering();
                var streamReader = new StreamReader(context.HttpContext.Request.Body);
                var parsedBody = JObject.Parse(streamReader.ReadToEnd());
                context.HttpContext.Request.Body.Position = 0;

                var validKeys = this.targetModelType.GetProperties().Select(getJsonName);
                var usedKeys = parsedBody.Children().Select(q => q.Path);

                foreach (var usedKey in usedKeys)
                {
                    if (usedKey == "id" || !validKeys.Contains(usedKey))
                    {
                        context.Result = new BadRequestObjectResult($"The provided data has an invalid attribute '{usedKey}'.");
                    }
                }
            }
            catch (JsonReaderException)
            {
                context.Result = new BadRequestObjectResult("Invalid JSON in body.");
            }
        }
    }
}
