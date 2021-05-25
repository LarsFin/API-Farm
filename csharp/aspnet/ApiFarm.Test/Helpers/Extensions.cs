using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Test.Helpers
{
    public static class Extensions
    {
        public static ObjectResult AsObjectResult<T>(this ActionResult<T> actionResult) => actionResult.Result as ObjectResult;
    }
}
