using System;
using System.IO;
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
                JObject.Parse(streamReader.ReadToEnd());
                context.HttpContext.Request.Body.Position = 0;
            }
            catch (JsonReaderException)
            {
                context.Result = new BadRequestObjectResult("Invalid JSON in body.");
            }
        }
    }
}
