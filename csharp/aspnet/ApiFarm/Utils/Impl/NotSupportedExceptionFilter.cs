using System;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// Filter to convey unsupported requests as bad ones as opposed to internal server faults.
    /// </summary>
    [AttributeUsage(AttributeTargets.Class)]
    public class NotSupportedExceptionFilter : ExceptionFilterAttribute
    {
        /// <summary>
        /// Called when an exception is thrown in process of a request. This filter's action is guarded in the instance the exception is not of
        /// type <see cref="NotSupportedException"/>.
        /// </summary>
        /// <param name="context">The context under which the exception was thrown.</param>
        public override void OnException(ExceptionContext context)
        {
            if (context.Exception.GetType() != typeof(NotSupportedException))
            {
                return;
            }

            context.Result = new BadRequestObjectResult(context.Exception.Message);
        }
    }
}
